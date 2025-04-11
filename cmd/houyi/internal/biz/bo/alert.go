package bo

import (
	"time"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
)

type Alert interface {
	GetStatus() common.EventStatus
	GetLabels() *label.Label
	GetAnnotations() *label.Annotation
	GetStartsAt() *time.Time
	GetEndsAt() *time.Time
	GetGeneratorURL() string
	GetFingerprint() string
	StatusNext() Alert
	IsFiring() bool
	GetValue() float64

	GetDuration() time.Duration
	IsSustaining() bool
}
