package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data"
)

func NewJudgeRepo(data *data.Data, logger log.Logger) repository.Judge {
	return &judgeImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.judge")),
	}
}

type judgeImpl struct {
	*data.Data
	helper *log.Helper
}

func (j *judgeImpl) Metric(ctx context.Context, data []bo.MetricJudgeData, rule bo.MetricJudgeRule) ([]bo.Alert, error) {
	//TODO implement me
	panic("implement me")
}
