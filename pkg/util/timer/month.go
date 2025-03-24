package timer

import (
	"time"
)

var _ Matcher = (*DayOfMonths)(nil)
var _ Matcher = (*Month)(nil)

type DayOfMonths struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func (m *DayOfMonths) Match(t time.Time) bool {
	d := t.Day()
	return d >= m.Start && d <= m.End
}

type Month struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func (m *Month) Match(t time.Time) bool {
	month := t.Month()
	if m.Start > m.End {
		return month >= time.Month(m.Start) || month <= time.Month(m.End)
	}
	return month >= time.Month(m.Start) && month <= time.Month(m.End)
}
