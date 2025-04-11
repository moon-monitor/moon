package event

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/robfig/cron/v3"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/repository"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/server"
)

func NewAlertJob(alert bo.Alert, opts ...AlertJobOption) (bo.AlertJob, error) {
	a := &alertJob{
		Alert: alert,
	}
	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}
	checkOpts := []*checkItem{
		{"alertRepo", a.alertRepo},
		{"eventBusRepo", a.eventBusRepo},
		{"helper", a.helper},
	}
	return a, checkList(checkOpts...)
}

type alertJob struct {
	bo.Alert

	id           cron.EntryID
	alertRepo    repository.Alert
	eventBusRepo repository.EventBus

	helper *log.Helper
}

type AlertJobOption func(*alertJob) error

func WithAlertJobAlertRepo(alertRepo repository.Alert) AlertJobOption {
	return func(a *alertJob) error {
		if alertRepo == nil {
			return merr.ErrorInternalServerError("alertRepo is nil")
		}
		a.alertRepo = alertRepo
		return nil
	}
}

func WithAlertJobEventBusRepo(eventBusRepo repository.EventBus) AlertJobOption {
	return func(a *alertJob) error {
		if eventBusRepo == nil {
			return merr.ErrorInternalServerError("eventBusRepo is nil")
		}
		a.eventBusRepo = eventBusRepo
		return nil
	}
}

func WithAlertJobHelper(logger log.Logger) AlertJobOption {
	return func(a *alertJob) error {
		if logger == nil {
			return merr.ErrorInternalServerError("logger is nil")
		}
		a.helper = log.NewHelper(log.With(logger, "module", "event.alert", "jobKey", a.GetFingerprint()))
		return nil
	}
}

func (a *alertJob) isSustaining() (alert bo.Alert, sustaining bool) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	defer func() {
		if sustaining {
			return
		}
		if err := a.alertRepo.Delete(ctx, a.GetFingerprint()); err != nil {
			a.helper.Warnw("msg", "delete alert error", "error", err)
		}
	}()
	alert, ok := a.alertRepo.Get(ctx, a.GetFingerprint())
	if !ok {
		return a, false
	}
	return alert, alert.GetLastUpdated().Add(a.GetDuration()).After(time.Now())
}

func (a *alertJob) Run() {
	alertInfo, ok := a.isSustaining()
	if !ok {
		alertInfo.Resolved()
		a.Alert = alertInfo
		a.eventBusRepo.InAlertEventBus() <- a
	}
}

func (a *alertJob) ID() cron.EntryID {
	if a == nil {
		return 0
	}
	return a.id
}

func (a *alertJob) Index() string {
	return a.GetFingerprint()
}

func (a *alertJob) Spec() server.CronSpec {
	if a == nil {
		return server.CronSpecEvery(1 * time.Minute)
	}
	return server.CronSpecEvery(a.GetDuration())
}

func (a *alertJob) WithID(id cron.EntryID) server.CronJob {
	a.id = id
	return a
}
