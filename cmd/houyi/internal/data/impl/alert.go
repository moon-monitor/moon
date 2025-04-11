package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/event"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/vobj"
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
	if len(alerts) == 0 {
		return nil
	}
	key := vobj.AlertEventCacheKey.Key()
	alertMap := make(map[string]any, len(alerts))
	for _, alert := range alerts {
		fingerprint := alert.GetFingerprint()
		item := &event.AlertJob{
			Status:       alert.GetStatus(),
			Labels:       alert.GetLabels(),
			Annotations:  alert.GetAnnotations(),
			StartsAt:     alert.GetStartsAt(),
			EndsAt:       alert.GetEndsAt(),
			GeneratorURL: alert.GetGeneratorURL(),
			Fingerprint:  fingerprint,
			LastUpdated:  time.Now(),
			Duration:     alert.GetDuration(),
		}
		alertMap[fingerprint] = item
	}
	if err := a.GetCache().Client().HSet(ctx, key, alertMap).Err(); err != nil {
		a.helper.Warnw("method", "SaveAlert", "err", err)
		return err
	}
	alertEventBus := a.InAlertEventBus()
	for _, alert := range alertMap {
		alertBo, ok := alert.(*event.AlertJob)
		if !ok {
			continue
		}
		alertEventBus <- alertBo
	}
	return nil
}
