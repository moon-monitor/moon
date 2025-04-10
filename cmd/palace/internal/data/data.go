package data

import (
	"database/sql"
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
	"github.com/moon-monitor/moon/pkg/util/safety"
)

// ProviderSetData is a set of data providers.
var ProviderSetData = wire.NewSet(New)

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
		rabbitConn: safety.NewMap[string, *bo.Server](),
		houyiConn:  safety.NewMap[string, *bo.Server](),
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
		for _, server := range data.rabbitConn.List() {
			if err = server.Conn.Close(); err != nil {
				data.helper.Errorw("method", "close rabbit conn", "err", err)
			}
		}
		for _, server := range data.houyiConn.List() {
			if err = server.Conn.Close(); err != nil {
				data.helper.Errorw("method", "close houyi conn", "err", err)
			}
		}
	}, nil
}

func newSqlDB(c *config.Database) (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?%s", c.GetUser(), c.GetPassword(), c.GetHost(), c.GetPort(), c.GetDbName(), c.GetParams())
	switch c.GetDriver() {
	case config.Database_MYSQL:
		return sql.Open("mysql", dsn)
	default:
		return nil, fmt.Errorf("unknown driver: %s", c.GetDriver())
	}
}

type Data struct {
	dataConf             *conf.Data
	mainDB               gorm.DB
	bizDB, eventDB       *sql.DB
	bizDBMap, eventDBMap *safety.Map[uint32, gorm.DB]
	cache                cache.Cache
	rabbitConn           *safety.Map[string, *bo.Server]
	houyiConn            *safety.Map[string, *bo.Server]
	helper               *log.Helper
}

func (d *Data) GetRabbitConn(id string) (*bo.Server, bool) {
	return d.rabbitConn.Get(id)
}

func (d *Data) SetRabbitConn(id string, conn *bo.Server) {
	d.rabbitConn.Set(id, conn)
}

func (d *Data) RemoveRabbitConn(id string) {
	d.rabbitConn.Delete(id)
}

func (d *Data) GetHouyiConn(id string) (*bo.Server, bool) {
	return d.houyiConn.Get(id)
}

func (d *Data) SetHouyiConn(id string, conn *bo.Server) {
	d.houyiConn.Set(id, conn)
}

func (d *Data) RemoveHouyiConn(id string) {
	d.houyiConn.Delete(id)
}

func (d *Data) FirstRabbitConn() (*bo.Server, bool) {
	list := d.rabbitConn.List()
	for _, conn := range list {
		return conn, true
	}
	return nil, false
}

func (d *Data) GetCache() cache.Cache {
	return d.cache
}

func (d *Data) GetMainDB() gorm.DB {
	return d.mainDB
}

func (d *Data) GetBizDB(teamID uint32) (gorm.DB, error) {
	db, ok := d.bizDBMap.Get(teamID)
	if ok {
		return db, nil
	}
	return nil, merr.ErrorInternalServerError("team db not found").WithMetadata(map[string]string{"method": "GetBizDB"})
}

func (d *Data) GetEventDB(teamID uint32) (gorm.DB, error) {
	db, ok := d.eventDBMap.Get(teamID)
	if ok {
		return db, nil
	}
	return nil, merr.ErrorInternalServerError("team db not found").WithMetadata(map[string]string{"method": "GetEventDB"})
}
