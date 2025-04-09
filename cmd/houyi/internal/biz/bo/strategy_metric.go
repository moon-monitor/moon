package bo

import (
	"time"

	"github.com/moon-monitor/moon/pkg/api/houyi/common"
)

type MetricStrategyItem struct {
	StrategyId     uint32
	Team           TeamItem
	Datasource     []MetricDatasourceConfig
	Name           string
	Expr           string
	ReceiverRoutes []string
	Labels         Label
	Annotations    Annotation
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
