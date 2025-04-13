package team

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameStrategyMetricRule = "team_strategy_metric_rules"

type StrategyMetricRule struct {
	do.BaseModel
	StrategyMetricID uint32                           `gorm:"column:strategy_metric_id;type:int unsigned;not null;comment:策略指标id" json:"strategyMetricID"`
	StrategyMetric   *StrategyMetric                  `gorm:"foreignKey:StrategyMetricID;references:ID" json:"strategyMetric"`
	LevelID          uint32                           `gorm:"column:level_id;type:int unsigned;not null;comment:等级id" json:"levelID"`
	Level            *Dict                            `gorm:"foreignKey:LevelID;references:ID" json:"level"`
	SampleMode       vobj.SampleMode                  `gorm:"column:sample_mode;type:tinyint(2);not null;comment:采样方式" json:"sampleMode"`
	Condition        vobj.ConditionMetric             `gorm:"column:condition;type:tinyint(2);not null;comment:条件" json:"condition"`
	Count            int64                            `gorm:"column:count;type:bigint;not null;comment:采样数量" json:"count"`
	Values           *Values                          `gorm:"column:values;type:json;not null;comment:值" json:"values"`
	Duration         time.Duration                    `gorm:"column:duration;type:time;not null;comment:持续时间" json:"duration"`
	Status           vobj.GlobalStatus                `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Notices          []*NoticeGroup                   `gorm:"many2many:team_strategy_metric_rule_notice_groups" json:"notices"`
	LabelNotices     []*StrategyMetricRuleLabelNotice `gorm:"foreignKey:StrategyMetricRuleID;references:ID" json:"labelNotices"`
	AlarmPages       []*Dict                          `gorm:"many2many:team_strategy_metric_rule_alarm_pages" json:"alarmPages"`
}

func (StrategyMetricRule) TableName() string {
	return tableNameStrategyMetricRule
}

var _ do.ORMModel = (*Values)(nil)

type Values []float64

func (v *Values) Scan(src any) error {
	switch value := src.(type) {
	case []byte:
		return json.Unmarshal(value, v)
	case string:
		return json.Unmarshal([]byte(value), v)
	default:
		return nil
	}
}

func (v Values) Value() (driver.Value, error) {
	return json.Marshal(v)
}
