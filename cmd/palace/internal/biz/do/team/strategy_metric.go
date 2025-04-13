package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

const tableNameStrategyMetrics = "team_strategy_metrics"

type StrategyMetric struct {
	do.BaseModel
	StrategyID          uint32                `gorm:"column:strategy_id;type:int unsigned;not null;comment:策略id" json:"strategyID"`
	Strategy            *Strategy             `gorm:"foreignKey:StrategyID;references:ID" json:"strategy"`
	Expr                string                `gorm:"column:expr;type:varchar(1024);not null;comment:表达式" json:"expr"`
	Labels              kv.StringMap          `gorm:"column:labels;type:json;not null;comment:标签" json:"labels"`
	Annotations         kv.StringMap          `gorm:"column:annotations;type:json;not null;comment:注解" json:"annotations"`
	StrategyMetricRules []*StrategyMetricRule `gorm:"foreignKey:StrategyMetricID;references:ID" json:"strategyMetricRules"`
	Datasource          []*DatasourceMetric   `gorm:"many2many:strategy_metric_datasource" json:"datasource"`
}

func (StrategyMetric) TableName() string {
	return tableNameStrategyMetrics
}
