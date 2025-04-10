package do

import (
	"encoding/json"
	"time"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/util/hash"
	"github.com/moon-monitor/moon/pkg/util/kv"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
)

var _ cache.Object = (*Alert)(nil)

type Alert struct {
	Status       common.EventStatus `json:"status"`
	Labels       *label.Label       `json:"labels"`
	Annotations  *label.Annotation  `json:"annotations"`
	StartsAt     *time.Time         `json:"startsAt"`
	EndsAt       *time.Time         `json:"endsAt"`
	GeneratorURL string             `json:"generatorURL"`
	Fingerprint  string             `json:"fingerprint"`
	Value        float64            `json:"value"`

	Duration    time.Duration `json:"duration"`
	LastUpdated time.Time     `json:"lastUpdated"`
}

func (a *Alert) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

func (a *Alert) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

func (a *Alert) UniqueKey() string {
	return a.GetFingerprint()
}

func (a *Alert) GetStatus() common.EventStatus {
	return a.Status
}

func (a *Alert) GetLabels() *label.Label {
	return a.Labels
}

func (a *Alert) GetAnnotations() *label.Annotation {
	return a.Annotations
}

func (a *Alert) GetStartsAt() *time.Time {
	return a.StartsAt
}

func (a *Alert) GetEndsAt() *time.Time {
	return a.EndsAt
}

func (a *Alert) GetGeneratorURL() string {
	return a.GeneratorURL
}

func (a *Alert) GetFingerprint() string {
	if a.Fingerprint == "" {
		stringMap := kv.NewStringMap(a.Labels.ToMap())
		a.Fingerprint = hash.MD5(kv.SortString(stringMap))
	}
	return a.Fingerprint
}

func (a *Alert) Resolved() {
	a.Status = common.EventStatus_resolved
	a.LastUpdated = time.Now()
	a.EndsAt = &a.LastUpdated
}

func (a *Alert) GetValue() float64 {
	return a.Value
}

func (a *Alert) GetDuration() time.Duration {
	return a.Duration
}

func (a *Alert) GetLastUpdated() time.Time {
	return a.LastUpdated
}

func (a *Alert) IsPending() bool {
	return a.Status == common.EventStatus_pending
}

func (a *Alert) IsFiring() bool {
	return a.Status == common.EventStatus_firing
}

func (a *Alert) IsResolved() bool {
	return a.Status == common.EventStatus_resolved
}

func (a *Alert) Firing() {
	a.Status = common.EventStatus_firing
	a.LastUpdated = time.Now()
	a.EndsAt = nil
}
