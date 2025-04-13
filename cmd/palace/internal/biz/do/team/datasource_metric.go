package team

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

const tableNameDatasourceMetric = "team_datasource_metrics"

type DatasourceMetric struct {
	do.TeamModel
	Name           string                      `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Status         vobj.GlobalStatus           `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Remark         string                      `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Driver         vobj.DatasourceDriverMetric `gorm:"column:type;type:tinyint(2);not null;comment:类型" json:"type"`
	Endpoint       string                      `gorm:"column:endpoint;type:varchar(255);not null;comment:数据源地址" json:"endpoint"`
	ScrapeInterval time.Duration               `gorm:"column:scrape_interval;type:time;not null;comment:抓取间隔" json:"scrapeInterval"`
	Headers        kv.StringMap                `gorm:"column:headers;type:text;not null;comment:请求头" json:"headers"`
	QueryMethod    vobj.HTTPMethod             `gorm:"column:query_method;type:tinyint(2);not null;comment:请求方法" json:"queryMethod"`
	CA             string                      `gorm:"column:ca;type:text;not null;comment:ca" json:"ca"`
	TLS            *do.TLS                     `gorm:"column:tls;type:text;not null;comment:tls" json:"tls"`
	BasicAuth      *do.BasicAuth               `gorm:"column:basic_auth;type:text;not null;comment:basic_auth" json:"basicAuth"`
	Extra          kv.StringMap                `gorm:"column:extra;type:text;not null;comment:额外信息" json:"extra"`
	Metrics        []*StrategyMetric           `gorm:"many2many:strategy_metric_datasource" json:"metrics"`
}

func (d *DatasourceMetric) TableName() string {
	return tableNameDatasourceMetric
}
