package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	apicommon "github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/util/kv/label"
	"github.com/moon-monitor/moon/pkg/util/timex"
)

func ToAlertParams(req *apicommon.AlertItem) *bo.Alert {
	annotations := label.NewAnnotationFromMap(req.GetAnnotations())
	labels := label.NewLabel(req.GetLabels())
	return &bo.Alert{
		Status:       vobj.AlertStatus(req.Status),
		Labels:       labels.ToMap(),
		Summary:      annotations.GetSummary(),
		Description:  annotations.GetDescription(),
		Value:        req.GetValue(),
		GeneratorURL: req.GetGeneratorURL(),
		TeamID:       labels.GetTeamId(),
		Fingerprint:  req.GetFingerprint(),
		StartsAt:     timex.ParseX(req.GetStartsAt()),
		EndsAt:       timex.ParseX(req.GetEndsAt()),
	}
}
