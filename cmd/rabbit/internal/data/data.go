package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/moon-monitor/moon/pkg/plugin/email"
	"github.com/moon-monitor/moon/pkg/util/safety"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/conf"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
)

// ProviderSetData is a set of data providers.
var ProviderSetData = wire.NewSet(New)

func New(c *conf.Bootstrap, logger log.Logger) (*Data, func(), error) {
	var err error
	dataConf := c.Data
	data := &Data{
		dataConf: dataConf,
		emails:   safety.NewMap[string, email.Email](),
		helper:   log.NewHelper(log.With(logger, "module", "data")),
	}
	data.cache, err = cache.NewCache(c.GetCache())
	if err != nil {
		return nil, nil, err
	}

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		if err = data.cache.Close(); err != nil {
			log.NewHelper(logger).Errorw("method", "close cache", "err", err)
		}
	}
	return data, cleanup, nil
}

type Data struct {
	dataConf *conf.Data
	cache    cache.Cache
	emails   *safety.Map[string, email.Email]

	helper *log.Helper
}

func (d *Data) GetCache() cache.Cache {
	return d.cache
}

func (d *Data) GetEmail(name string) (email.Email, bool) {
	return d.emails.Get(name)
}

func (d *Data) SetEmail(name string, email email.Email) {
	d.emails.Set(name, email)
}
