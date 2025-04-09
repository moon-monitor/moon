package bo

import (
	"time"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
)

type MetricStrategyItem struct {
	StrategyId     uint32
	Team           TeamItem
	Datasource     []MetricDatasourceConfig
	Name           string
	Expr           string
	ReceiverRoutes []string
	Labels         *label.Label
	Annotations    *label.Annotation
	Duration       time.Duration
	Rules          []*MetricRuleItem
}

type MetricRuleItem struct {
	StrategyId     uint32
	LevelId        uint32
	LevelName      string
	Count          int64
	Values         []float64
	SampleMode     common.MetricStrategyItem_SampleMode
	Condition      common.MetricStrategyItem_Condition
	ReceiverRoutes []string
	LabelNotices   []*LabelNotices
}

type MetricRule interface {
	GetTeamId() uint32
	GetExpr() string
	GetDatasource() MetricDatasourceConfig
	GetStrategyId() uint32
	GetLevelId() uint32
	GetLabels() *label.Label
	GetAnnotations() *label.Annotation
	GetCount() int64
	GetValues() []float64
	GetSampleMode() common.MetricStrategyItem_SampleMode
	GetCondition() common.MetricStrategyItem_Condition
	GetReceiverRoutes() []string
	GetLabelNotices() []*LabelNotices
}
