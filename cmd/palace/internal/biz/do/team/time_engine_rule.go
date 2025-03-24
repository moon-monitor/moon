package team

import (
	"database/sql/driver"
	"encoding/json"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

const tableNameTimeEngineRule = "team_time_engine_rules"

type TimeEngineRule struct {
	do.TeamModel

	Name    string                    `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark  string                    `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status  vobj.TimeEngineRuleStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Rule    Rules                     `gorm:"column:rule;type:text;not null;comment:规则" json:"rule"`
	Type    vobj.TimeEngineRuleType   `gorm:"column:type;type:tinyint(2);not null;comment:类型" json:"type"`
	Engines []*TimeEngine             `gorm:"many2many:team_time_engine__time_rules" json:"engines"`
}

func (t *TimeEngineRule) TableName() string {
	return tableNameTimeEngineRule
}

type Rules []int

func (r Rules) Value() (driver.Value, error) {
	return json.Marshal(r)
}

func (r *Rules) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), r)
}
