package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/repository"
)

func NewHealthBiz(cacheRepo repository.Cache, logger log.Logger) *HealthBiz {
	return &HealthBiz{
		cacheRepo: cacheRepo,
		helper:    log.NewHelper(log.With(logger, "module", "biz.health")),
	}
}

type HealthBiz struct {
	cacheRepo repository.Cache

	helper *log.Helper
}
