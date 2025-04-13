package team

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameDatasourceMetric = "team_datasource_metrics"

type DatasourceMetric struct {
	do.TeamModel
	Name     string                      `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Driver   vobj.DatasourceDriverMetric `gorm:"column:type;type:tinyint(2);not null;comment:类型" json:"type"`
	Endpoint string                      `gorm:"column:endpoint;type:varchar(255);not null;comment:数据源地址" json:"endpoint"`

	Metrics []*StrategyMetric `gorm:"many2many:strategy_metric_datasource" json:"metrics"`
}

func (d *DatasourceMetric) TableName() string {
	return tableNameDatasourceMetric
}
