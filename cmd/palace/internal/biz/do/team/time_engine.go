package team

import (
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/timer"
)

const tableNameTimeEngine = "team_time_engines"

type TimeEngine struct {
	do.TeamModel
	Name   string            `gorm:"column:name;type:varchar(64);not null;comment:名称" json:"name"`
	Remark string            `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`
	Status vobj.GlobalStatus `gorm:"column:status;type:tinyint(2);not null;comment:状态" json:"status"`
	Rules  []*TimeEngineRule `gorm:"many2many:team_time_engine__time_rules" json:"rules"`
}

func (t *TimeEngine) TableName() string {
	return tableNameTimeEngine
}

func (t *TimeEngine) Allow(g time.Time) (bool, error) {
	matchers := make([]timer.Matcher, 0, len(t.Rules))
	errs := make([]error, 0, len(t.Rules))
	for _, rule := range t.Rules {
		matcher, err := rule.ToTimerMatcher()
		if err != nil {
			errs = append(errs, err)
			continue
		}
		matchers = append(matchers, matcher)
	}
	if len(errs) > 0 {
		return false, merr.ErrorParamsError("failed to convert time engine rules to timer matchers: %v", errs)
	}
	for _, matcher := range matchers {
		if !matcher.Match(g) {
			return false, nil
		}
	}
	return true, nil
}
