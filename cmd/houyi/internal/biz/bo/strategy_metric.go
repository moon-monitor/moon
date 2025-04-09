package bo

import (
	"time"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
)

type MetricJudgeRule interface {
	GetLabels() *label.Label
	GetAnnotations() *label.Annotation
	GetDuration() time.Duration
	GetCount() int64
	GetValues() []float64
	GetSampleMode() common.MetricStrategyItem_SampleMode
	GetCondition() common.MetricStrategyItem_Condition
}

type MetricJudgeDataValue interface {
	GetValue() float64
	GetTimestamp() int64
}

type MetricJudgeData interface {
	GetLabels() map[string]string
	GetValues() []MetricJudgeDataValue
}

type MetricRule interface {
	cache.Object
	GetTeamId() uint32
	GetDatasource() string
	GetStrategyId() uint32
	GetLevelId() uint32
	GetReceiverRoutes() []string
	GetLabelReceiverRoutes() []LabelNotices
	GetExpr() string

	MetricJudgeRule
}
