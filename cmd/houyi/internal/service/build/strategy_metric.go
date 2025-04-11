package build

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/api/houyi/common"
	"github.com/moon-monitor/moon/pkg/util/cnst"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func ToMetricRules(strategyItems ...*common.MetricStrategyItem) []bo.MetricRule {
	if len(strategyItems) == 0 {
		return nil
	}
	rules := make([]bo.MetricRule, 0, len(strategyItems)*5*3)
	for _, strategyItem := range strategyItems {
		if strategyItem == nil {
			continue
		}
		for _, rule := range strategyItem.Rules {
			if rule == nil {
				continue
			}
			datasourceConfigs := strategyItem.GetDatasource()
			for _, datasourceItem := range datasourceConfigs {
				if datasourceItem == nil {
					continue
				}
				annotations := strategyItem.GetAnnotations()
				item := &do.MetricRule{
					TeamId:     strategyItem.GetTeam().GetTeamId(),
					Datasource: vobj.MetricDatasourceUniqueKey(datasourceItem.GetDriver(), datasourceItem.GetId()),
					StrategyId: strategyItem.GetStrategyId(),
					LevelId:    rule.GetLevelId(),
					Receiver:   rule.GetReceiverRoutes(),
					LabelReceiver: slices.Map(rule.GetLabelNotices(), func(item *common.MetricStrategyItem_LabelNotices) *do.LabelNotices {
						return ToLabelNotice(item)
					}),
					Expr:        strategyItem.GetExpr(),
					Labels:      label.NewLabel(strategyItem.GetLabels()),
					Annotations: label.NewAnnotation(annotations[cnst.AnnotationKeySummary], annotations[cnst.AnnotationKeyDescription]),
					Duration:    rule.GetDuration().AsDuration(),
					Count:       rule.GetCount(),
					Values:      rule.GetValues(),
					SampleMode:  rule.GetSampleMode(),
					Condition:   rule.GetCondition(),
					Enable:      rule.GetEnable(),
				}
				rules = append(rules, item)
			}
		}
	}
	return rules
}
