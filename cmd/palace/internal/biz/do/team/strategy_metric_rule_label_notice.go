package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

const tableNameStrategyMetricRuleLabelNotice = "team_strategy_metric_rule_label_notices"

type StrategyMetricRuleLabelNotice struct {
	do.BaseModel
	StrategyMetricRuleID uint32         `gorm:"column:strategy_metric_rule_id;not null;int(10) unsigned;index:idx_strategy_metric_rule_id;comment:策略指标规则ID"`
	LabelKey             string         `gorm:"column:label_key;type:varchar(64);not null;comment:标签键" json:"labelKey"`
	LabelValue           string         `gorm:"column:label_value;type:varchar(255);not null;comment:标签值" json:"labelValue"`
	Notices              []*NoticeGroup `gorm:"many2many:team_strategy_metric_rule_label_notice_notice_groups" json:"notices"`
}

func (s *StrategyMetricRuleLabelNotice) TableName() string {
	return tableNameStrategyMetricRuleLabelNotice
}
