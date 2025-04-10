package event

import (
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
)

func NewAlertJob(opts ...AlertJobOption) (*AlertJob, error) {
	a := &AlertJob{}
	for _, opt := range opts {
		if err := opt(a); err != nil {
			return nil, err
		}
	}
	return a, nil
}

type AlertJob struct {
	Status       common.EventStatus
	Labels       *label.Label
	Annotations  *label.Annotation
	StartsAt     *time.Time
	EndsAt       *time.Time
	GeneratorURL string
	Fingerprint  string
}

func (a *AlertJob) GetStatus() common.EventStatus {
	if a == nil {
		return common.EventStatus_pending
	}
	return a.Status
}

func (a *AlertJob) GetLabels() *label.Label {
	if a == nil {
		return nil
	}
	return a.Labels
}

func (a *AlertJob) GetAnnotations() *label.Annotation {
	if a == nil {
		return nil
	}
	return a.Annotations
}

func (a *AlertJob) GetStartsAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.StartsAt
}

func (a *AlertJob) GetEndsAt() *time.Time {
	if a == nil {
		return nil
	}
	return a.EndsAt
}

func (a *AlertJob) GetGeneratorURL() string {
	if a == nil {
		return ""
	}
	return a.GeneratorURL
}

func (a *AlertJob) GetFingerprint() string {
	if a == nil {
		return ""
	}
	return a.Fingerprint
}

func (a *AlertJob) StatusNext() bo.Alert {
	switch a.Status {
	case common.EventStatus_firing:
		a.Status = common.EventStatus_resolved
	case common.EventStatus_pending:
		a.Status = common.EventStatus_firing
	default:
		a.Status = common.EventStatus_resolved
	}
	return a
}

func (a *AlertJob) IsFiring() bool {
	return a.Status == common.EventStatus_firing
}

type AlertJobOption func(*AlertJob) error
