package data

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	ggorm "gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/email"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
	"github.com/moon-monitor/moon/pkg/util/safety"
)

// ProviderSetData is a set of data providers.
var ProviderSetData = wire.NewSet(New)

type Data struct {
	dataConf             *conf.Data
	mainDB               gorm.DB
	bizDB, eventDB       *sql.DB
	bizDBMap, eventDBMap *safety.Map[uint32, gorm.DB]
	cache                cache.Cache
	email                email.Email

	helper *log.Helper
}

func (d *Data) GetCache() cache.Cache {
	return d.cache
}

func (d *Data) GetEmail() email.Email {
	return d.email
}

func (d *Data) GetMainDB() gorm.DB {
	return d.mainDB
}

func (d *Data) GetBizDB(teamID uint32) (gorm.DB, error) {
	db, ok := d.bizDBMap.Get(teamID)
	if ok {
		return db, nil
	}

	bizConf := d.dataConf.GetBiz()
	gormDB, err := d.createDatabase(bizConf, d.bizDB, teamID)
	if err != nil {
		return nil, err
	}
	d.bizDBMap.Set(teamID, gormDB)
	return d.GetBizDB(teamID)
}

func (d *Data) GetEventDB(teamID uint32) (gorm.DB, error) {
	db, ok := d.eventDBMap.Get(teamID)
	if ok {
		return db, nil
	}
	eventConf := d.dataConf.GetAlarm()
	gormDB, err := d.createDatabase(eventConf, d.eventDB, teamID)
	if err != nil {
		return nil, err
	}
	d.eventDBMap.Set(teamID, gormDB)
	return d.GetEventDB(teamID)
}

func (d *Data) createDatabase(c *conf.Data_Database, sqlDB *sql.DB, teamID uint32) (gorm.DB, error) {
	teamQuery := systemQuery.Use(d.GetMainDB().GetDB()).Team
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	teamDo, err := teamQuery.WithContext(ctx).Where(teamQuery.ID.Eq(teamID)).First()
	if err != nil {
		if errors.Is(err, ggorm.ErrRecordNotFound) {
			return nil, merr.ErrorTeamNotFound("team %d not found", teamID)
		}
		return nil, err
	}

	dbName := c.GetDbName()
	if teamDo.Capacity.AllowGroup() && c.IsGroup() {
		dbName = fmt.Sprintf("%s_%d", dbName, teamID)
	}
	dsn := c.GenDsn(dbName)
	switch c.GetDriver() {
	case config.Database_MYSQL:
		expr := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbName)
		if _, err := sqlDB.Exec(expr); err != nil {
			return nil, err
		}
	}

	databaseConf := &config.Database{
		Driver:       c.GetDriver(),
		Dsn:          dsn,
		Debug:        c.GetDebug(),
		UseSystemLog: c.GetUseSystemLog(),
	}
	return gorm.NewDB(databaseConf)
}

// New a data and returns.
func New(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	var err error
	data := &Data{
		dataConf:   c.GetData(),
		mainDB:     nil,
		bizDB:      nil,
		eventDB:    nil,
		bizDBMap:   safety.NewMap[uint32, gorm.DB](),
		eventDBMap: safety.NewMap[uint32, gorm.DB](),
		cache:      nil,
		email:      email.New(c.GetEmail()),
		helper:     log.NewHelper(log.With(logger, "module", "data")),
	}

	dataConf := c.GetData()
	data.bizDB, err = newSqlDB(dataConf.GetBiz())
	if err != nil {
		return nil, nil, err
	}
	data.eventDB, err = newSqlDB(dataConf.GetAlarm())
	if err != nil {
		return nil, nil, err
	}

	data.mainDB, err = gorm.NewDB(dataConf.GetMain())
	if err != nil {
		return nil, nil, err
	}
	data.cache, err = cache.NewCache(c.GetCache())
	if err != nil {
		return nil, nil, err
	}

	return data, func() {
		data.helper.Info("closing the data resources")
		if err = data.mainDB.Close(); err != nil {
			data.helper.Errorw("method", "close main db", "err", err)
		}
		if err = data.cache.Close(); err != nil {
			data.helper.Errorw("method", "close cache", "err", err)
		}
		for teamID, db := range data.bizDBMap.List() {
			if err = db.Close(); err != nil {
				method := fmt.Sprintf("close team [%d] biz db", teamID)
				data.helper.Errorw("method", method, "err", err)
			}
		}
		for teamID, db := range data.eventDBMap.List() {
			if err = db.Close(); err != nil {
				method := fmt.Sprintf("close team [%d] alarm db", teamID)
				data.helper.Errorw("method", method, "err", err)
			}
		}
		if err = data.bizDB.Close(); err != nil {
			data.helper.Errorw("method", "close bizDB", "err", err)
		}
		if err = data.eventDB.Close(); err != nil {
			data.helper.Errorw("method", "close eventDB", "err", err)
		}
	}, nil
}

func newSqlDB(c *conf.Data_Database) (*sql.DB, error) {
	dsn := c.GenDsn("")
	switch c.GetDriver() {
	case config.Database_MYSQL:
		return sql.Open("mysql", dsn)
	case config.Database_SQLITE:
		return sql.Open("sqlite3", dsn)
	default:
		return nil, fmt.Errorf("unknown driver: %s", c.GetDriver())
	}
}
