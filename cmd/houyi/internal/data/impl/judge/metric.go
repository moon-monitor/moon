package judge

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/data/impl/judge/condition"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
)

func NewMetricJudge(sampleModeType common.MetricStrategyItem_SampleMode, opts ...MetricJudgeOption) MetricJudge {
	var config MetricJudgeConfig
	for _, opt := range opts {
		opt(&config)
	}
	switch sampleModeType {
	case common.MetricStrategyItem_SampleMode_For:
		return &metricForJudge{config}
	case common.MetricStrategyItem_SampleMode_Max:
		return &metricMaxJudge{config}
	case common.MetricStrategyItem_SampleMode_Min:
		return &metricMinJudge{config}
	default:
		return &metricForJudge{config}
	}
}

type MetricJudgeConfig struct {
	condition       condition.MetricCondition
	conditionValues []float64
	conditionCount  int64
}

type MetricJudgeOption func(*MetricJudgeConfig)

func WithMetricJudgeCondition(condition condition.MetricCondition) MetricJudgeOption {
	return func(c *MetricJudgeConfig) {
		c.condition = condition
	}
}

func WithMetricJudgeConditionValues(values []float64) MetricJudgeOption {
	return func(c *MetricJudgeConfig) {
		c.conditionValues = values
	}
}

func WithMetricJudgeConditionCount(count int64) MetricJudgeOption {
	return func(c *MetricJudgeConfig) {
		c.conditionCount = count
	}
}

type MetricJudge interface {
	Judge(originValues []bo.MetricJudgeDataValue) (bo.MetricJudgeDataValue, bool)
	Type() common.MetricStrategyItem_SampleMode
}

type metricForJudge struct {
	MetricJudgeConfig
}

func (m *metricForJudge) Type() common.MetricStrategyItem_SampleMode {
	return common.MetricStrategyItem_SampleMode_For
}

func (m *metricForJudge) Judge(originValues []bo.MetricJudgeDataValue) (bo.MetricJudgeDataValue, bool) {
	total := int64(0)
	var currentValue bo.MetricJudgeDataValue
	for _, value := range originValues {
		if !m.condition.Comparable(m.conditionValues, value.GetValue()) {
			total = 0
			currentValue = nil
			continue
		}
		total += 1
		currentValue = value
		if total >= m.conditionCount {
			return currentValue, true
		}
	}
	return currentValue, total >= m.conditionCount && currentValue != nil
}

type metricMaxJudge struct {
	MetricJudgeConfig
}

func (m *metricMaxJudge) Type() common.MetricStrategyItem_SampleMode {
	return common.MetricStrategyItem_SampleMode_Max
}

func (m *metricMaxJudge) Judge(originValues []bo.MetricJudgeDataValue) (bo.MetricJudgeDataValue, bool) {
	total := int64(0)
	var currentValue bo.MetricJudgeDataValue
	for _, value := range originValues {
		if !m.condition.Comparable(m.conditionValues, value.GetValue()) {
			currentValue = value
			continue
		}

		total += 1
		if total > m.conditionCount {
			return nil, false
		}
	}
	return currentValue, currentValue != nil
}

type metricMinJudge struct {
	MetricJudgeConfig
}

func (m *metricMinJudge) Type() common.MetricStrategyItem_SampleMode {
	return common.MetricStrategyItem_SampleMode_Min
}

func (m *metricMinJudge) Judge(originValues []bo.MetricJudgeDataValue) (bo.MetricJudgeDataValue, bool) {
	total := int64(0)
	var currentValue bo.MetricJudgeDataValue
	for _, value := range originValues {
		if m.condition.Comparable(m.conditionValues, value.GetValue()) {
			total += 1
			currentValue = value
			if total >= m.conditionCount {
				return currentValue, true
			}
		}
	}
	return nil, false
}
