package data

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/email"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
	"github.com/moon-monitor/moon/pkg/util/safety"
)

// ProviderSetData is a set of data providers.
var ProviderSetData = wire.NewSet(New)

type Data struct {
	dataConf       *conf.Data
	mainDB         gorm.DB
	bizDB, eventDB *safety.Map[uint32, gorm.DB]
	cache          cache.Cache
	email          email.Email

	helper *log.Helper
}

func (d *Data) GetMainDB() gorm.DB {
	return d.mainDB
}

func (d *Data) GetBizDB(teamID uint32) (gorm.DB, error) {
	db, ok := d.bizDB.Get(teamID)
	if !ok {
		return db, nil
	}
	bizConf := d.dataConf.GetBiz()
	c := &config.Database{
		Driver:       bizConf.GetDriver(),
		Dsn:          fmt.Sprintf(bizConf.GetDsn(), teamID),
		Debug:        bizConf.GetDebug(),
		UseSystemLog: bizConf.GetUseSystemLog(),
	}
	gormDB, err := gorm.NewDB(c)
	if err != nil {
		return nil, err
	}
	d.bizDB.Set(teamID, gormDB)
	return d.GetBizDB(teamID)
}

func (d *Data) GetEventDB(teamID uint32) (gorm.DB, error) {
	db, ok := d.eventDB.Get(teamID)
	if ok {
		return db, nil
	}
	eventConf := d.dataConf.GetAlarm()
	c := &config.Database{
		Driver:       eventConf.GetDriver(),
		Dsn:          fmt.Sprintf(eventConf.GetDsn(), teamID),
		Debug:        eventConf.GetDebug(),
		UseSystemLog: eventConf.GetUseSystemLog(),
	}
	gormDB, err := gorm.NewDB(c)
	if err != nil {
		return nil, err
	}
	d.bizDB.Set(teamID, gormDB)
	return d.GetEventDB(teamID)
}

func (d *Data) GetCache() cache.Cache {
	return d.cache
}

func (d *Data) GetEmail() email.Email {
	return d.email
}

// New a data and returns.
func New(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	var err error
	data := &Data{
		dataConf: c.GetData(),
		mainDB:   nil,
		bizDB:    safety.NewMap[uint32, gorm.DB](),
		eventDB:  safety.NewMap[uint32, gorm.DB](),
		cache:    nil,
		email:    email.New(c.GetEmail()),
		helper:   log.NewHelper(log.With(logger, "module", "data")),
	}

	dataConf := c.GetData()
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
		for teamID, db := range data.bizDB.List() {
			if err = db.Close(); err != nil {
				method := fmt.Sprintf("close team [%d] biz db", teamID)
				data.helper.Errorw("method", method, "err", err)
			}
		}
		for teamID, db := range data.eventDB.List() {
			if err = db.Close(); err != nil {
				method := fmt.Sprintf("close team [%d] alarm db", teamID)
				data.helper.Errorw("method", method, "err", err)
			}
		}
	}, nil
}
