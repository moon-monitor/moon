package event

import (
	"encoding/json"
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/util/hash"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
	"github.com/moon-monitor/moon/pkg/util/pointer"
)

var _ cache.Object = (*AlertJob)(nil)

type AlertJob struct {
	Status       common.EventStatus `json:"status"`
	Labels       *label.Label       `json:"labels"`
	Annotations  *label.Annotation  `json:"annotations"`
	StartsAt     *time.Time         `json:"startsAt"`
	EndsAt       *time.Time         `json:"endsAt"`
	GeneratorURL string             `json:"generatorURL"`
	Fingerprint  string             `json:"fingerprint"`
	Value        float64            `json:"value"`

	LastUpdated time.Time     `json:"lastUpdated"`
	Duration    time.Duration `json:"duration"`
}

func (a *AlertJob) GetValue() float64 {
	return a.Value
}

func (a *AlertJob) GetDuration() time.Duration {
	return a.Duration
}

func (a *AlertJob) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *AlertJob) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a *AlertJob) UniqueKey() string {
	return a.GetFingerprint()
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
	if a.Fingerprint == "" {
		return a.Fingerprint
	}
	stringMap := kv.NewStringMap(a.Labels.ToMap())
	a.Fingerprint = hash.MD5(kv.SortString(stringMap))
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
	if a.Status == common.EventStatus_resolved && a.EndsAt == nil {
		a.EndsAt = pointer.Of(time.Now().UTC())
	}
	return a
}

func (a *AlertJob) IsFiring() bool {
	return a.Status == common.EventStatus_firing
}

func (a *AlertJob) IsSustaining() bool {
	return time.Now().Add(-a.Duration).Before(a.LastUpdated)
}
