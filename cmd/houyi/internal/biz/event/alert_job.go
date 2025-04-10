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
}

func (a *AlertJob) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *AlertJob) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a *AlertJob) UniqueKey() string {
	return a.Fingerprint
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
	stringMap := kv.NewStringMap(a.Labels.ToMap())
	return hash.MD5(kv.SortString(stringMap))
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
