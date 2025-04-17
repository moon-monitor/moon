package do

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

type StrategyGroup interface {
	TeamBase
	GetName() string
	GetRemark() string
	GetStatus() vobj.GlobalStatus
	GetStrategies() []Strategy
}

type Strategy interface {
	TeamBase
	GetStrategyGroupID() uint32
	GetStrategyGroup() StrategyGroup
	GetStatus() vobj.GlobalStatus
	GetName() string
	GetRemark() string
	GetNotices() []NoticeGroup
}

type StrategyMetric interface {
	TeamBase
	GetStrategyID() uint32
	GetStrategy() Strategy
	GetDatasourceList() []DatasourceMetric
	GetExpr() string
	GetLabels() kv.StringMap
	GetAnnotations() kv.StringMap
	GetRules() []StrategyMetricRule
}

type StrategyMetricRule interface {
	TeamBase
	GetStrategyMetricID() uint32
	GetStrategyMetric() StrategyMetric
	GetLevelID() uint32
	GetLevel() Dict
	GetSampleMode() vobj.SampleMode
	GetCondition() vobj.ConditionMetric
	GetCount() int64
	GetValues() []float64
	GetDuration() time.Duration
	GetStatus() vobj.GlobalStatus
	GetNotices() []NoticeGroup
	GetLabelNotices() []StrategyMetricRuleLabelNotice
	GetAlarmPages() []Dict
}

type StrategyMetricRuleLabelNotice interface {
	TeamBase
	GetStrategyMetricRuleID() uint32
	GetLabelKey() string
	GetLabelValue() string
	GetNotices() []NoticeGroup
}
