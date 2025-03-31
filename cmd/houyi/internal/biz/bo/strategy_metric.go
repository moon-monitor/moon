package bo

import (
	"time"

	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/plugin/datasource/prometheus"
	"github.com/moon-monitor/moon/pkg/plugin/datasource/victoria"
)

type MetricDatasourceItem struct {
	Team            TeamItem
	Driver          vobj.MetricDatasourceDriver
	Prometheus      prometheus.Config
	VictoriaMetrics victoria.Config
}

type MetricStrategyItem struct {
	StrategyId     uint32
	Team           TeamItem
	Datasource     []*MetricDatasourceItem
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
	SampleMode     vobj.MetricStrategySampleMode
	Condition      vobj.MetricStrategyCondition
	ReceiverRoutes []string
	LabelNotices   []*LabelNotices
}
