package impl

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
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

func (a *alertImpl) Get(ctx context.Context, fingerprint string) (bo.Alert, bool) {
	key := vobj.AlertEventCacheKey.Key()
	exist, err := a.GetCache().Client().HExists(ctx, key, fingerprint).Result()
	if err != nil {
		a.helper.Warnw("method", "GetAlert", "err", err)
		return nil, false
	}
	if !exist {
		return nil, false
	}
	var alert do.Alert
	if err := a.GetCache().Client().HGet(ctx, key, fingerprint).Scan(&alert); err != nil {
		a.helper.Warnw("method", "GetAlert", "err", err)
		return nil, false
	}
	return &alert, true
}

func (a *alertImpl) Save(ctx context.Context, alerts ...bo.Alert) error {
	if len(alerts) == 0 {
		return nil
	}
	key := vobj.AlertEventCacheKey.Key()
	alertMap := make(map[string]any, len(alerts))
	for _, alert := range alerts {
		fingerprint := alert.GetFingerprint()
		item := &do.Alert{
			Status:       alert.GetStatus(),
			Labels:       alert.GetLabels(),
			Annotations:  alert.GetAnnotations(),
			StartsAt:     alert.GetStartsAt(),
			EndsAt:       alert.GetEndsAt(),
			GeneratorURL: alert.GetGeneratorURL(),
			Fingerprint:  fingerprint,
			Value:        alert.GetValue(),
			Duration:     alert.GetDuration(),
			LastUpdated:  time.Now(),
		}
		alertMap[fingerprint] = item
	}
	if err := a.GetCache().Client().HSet(ctx, key, alertMap).Err(); err != nil {
		a.helper.Warnw("method", "SaveAlert", "err", err)
		return err
	}
	return nil
}
