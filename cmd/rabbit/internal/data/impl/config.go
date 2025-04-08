package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/data"
)

func NewConfigRepo(d *data.Data, logger log.Logger) repository.Config {
	return &configImpl{
		helper: log.NewHelper(log.With(logger, "module", "data.repo.config")),
		Data:   d,
	}
}

type configImpl struct {
	helper *log.Helper
	*data.Data
}

func (c *configImpl) GetEmailConfig(ctx context.Context, name string) (bo.EmailConfig, bool) {
	var emailConfig do.EmailConfig
	key := vobj.EmailCacheKey.Key(name)
	result, err := c.Data.GetCache().Client().Exists(ctx, key).Result()
	if err != nil {
		c.helper.Errorw("GetEmailConfig", "err", err)
		return nil, false
	}
	if result == 0 {
		return nil, false
	}
	if err := c.Data.GetCache().Client().Get(ctx, key).Scan(&emailConfig); err != nil {
		c.helper.Errorw("GetEmailConfig", "err", err)
		return nil, false
	}
	return &emailConfig, true
}
