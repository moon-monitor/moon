package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data"
)

func NewAlertRepo(data *data.Data, logger log.Logger) repository.Alert {
	return &alertImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.alert")),
	}
}

type alertImpl struct {
	*data.Data
	helper *log.Helper
}

func (a *alertImpl) Save(ctx context.Context, alerts ...bo.Alert) error {
	//TODO implement me
	panic("implement me")
}
