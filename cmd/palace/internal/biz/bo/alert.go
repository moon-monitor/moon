package bo

import (
	"strings"
	"time"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/kv"
)

type Alert struct {
	TeamID       uint32           `json:"teamId"`
	Status       vobj.AlertStatus `json:"status"`
	Fingerprint  string           `json:"fingerprint"`
	Labels       kv.StringMap     `json:"labels"`
	Summary      string           `json:"summary"`
	Description  string           `json:"description"`
	Value        string           `json:"value"`
	GeneratorURL string           `json:"generatorURL"`
	StartsAt     time.Time        `json:"startsAt"`
	EndsAt       time.Time        `json:"endsAt"`
}

func (a *Alert) Validate() error {
	if a.StartsAt.IsZero() {
		return merr.ErrorParamsError("startsAt is required")
	}
	if !a.Status.Exist() {
		return merr.ErrorParamsError("status is required")
	}
	if strings.TrimSpace(a.Fingerprint) == "" {
		return merr.ErrorParamsError("fingerprint is required")
	}
	if a.TeamID <= 0 {
		return merr.ErrorParamsError("teamId is required")
	}

	return nil
}

type GetAlertParams struct {
	TeamID      uint32    `json:"teamId"`
	Fingerprint string    `json:"fingerprint"`
	StartsAt    time.Time `json:"startsAt"`
}
