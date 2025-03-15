package data

import (
	"fmt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
	"github.com/moon-monitor/moon/pkg/util/safety"
)

// ProviderSetData is a set of data providers.
var ProviderSetData = wire.NewSet(New)

type Data struct {
	mainDB         gorm.DB
	bizDB, alarmDB safety.Map[uint32, gorm.DB]
	cache          cache.Cache

	helper *log.Helper
}

// New a data and returns.
func New(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	var (
		data Data
		err  error
	)
	data.helper = log.NewHelper(log.With(logger, "module", "data"))
	dataConf := c.GetData()
	data.mainDB, err = gorm.NewDB(dataConf.GetMain())
	if err != nil {
		return nil, nil, err
	}
	data.cache, err = cache.NewCache(c.GetCache())
	if err != nil {
		return nil, nil, err
	}

	// TODO get all team, to init bizDB, alarmDB

	return &data, func() {
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
		for teamID, db := range data.alarmDB.List() {
			if err = db.Close(); err != nil {
				method := fmt.Sprintf("close team [%d] alarm db", teamID)
				data.helper.Errorw("method", method, "err", err)
			}
		}
	}, nil
}
