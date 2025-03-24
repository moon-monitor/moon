package timer

import (
	"slices"
	"time"
)

var _ Matcher = (*DaysOfWeek)(nil)

type DaysOfWeek struct {
	Days []time.Weekday `json:"days"`
}

func (d *DaysOfWeek) Match(t time.Time) bool {
	return slices.Contains(d.Days, t.Weekday())
}
