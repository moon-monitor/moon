package timer

import (
	"slices"
	"time"

	"github.com/moon-monitor/moon/pkg/merr"
)

var _ Matcher = (*HourRange)(nil)
var _ Matcher = (*Hour)(nil)

func NewHourRange(rule []int) (*HourRange, error) {
	if len(rule) != 2 {
		return nil, merr.ErrorParamsError("invalid hour range: %v", rule)
	}

	start := rule[0]
	end := rule[1]
	if start < 0 || start > 23 || end < 0 || end > 23 {
		return nil, merr.ErrorParamsError("invalid hour range: %d-%d", start, end)
	}
	return &HourRange{Start: start, End: end}, nil
}

type HourRange struct {
	Start int `json:"start"`
	End   int `json:"end"`
}

func (h *HourRange) Match(t time.Time) bool {
	hour := t.Hour()
	if h.Start > h.End {
		return hour >= h.Start || hour <= h.End
	}
	return hour >= h.Start && hour <= h.End
}

func NewHour(rule []int) (*Hour, error) {
	return &Hour{Hours: rule}, nil
}

type Hour struct {
	Hours []int `json:"hours"`
}

func (h *Hour) Match(t time.Time) bool {
	hour := t.Hour()
	return slices.Contains(h.Hours, hour)
}

type HourMinute struct {
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

func (h *HourMinute) GT(t time.Time) bool {
	hour := t.Hour()
	minute := t.Minute()
	return h.Hour > hour || (h.Hour == hour && h.Minute > minute)
}

type HourMinuteRange struct {
	Start HourMinute `json:"start"`
	End   HourMinute `json:"end"`
}
