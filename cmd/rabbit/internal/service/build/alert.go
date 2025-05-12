package build

import (
	"github.com/moon-monitor/moon/cmd/rabbit/internal/biz/bo"
	"github.com/moon-monitor/moon/pkg/api/common"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func ToAlerts(req *common.AlertsItem) *bo.AlertsItem {
	return &bo.AlertsItem{
		Receiver:          req.GetReceiver(),
		Status:            req.GetStatus().String(),
		Alerts:            ToAlertItems(req.GetAlerts()),
		GroupLabels:       req.GetGroupLabels(),
		CommonLabels:      req.GetCommonLabels(),
		CommonAnnotations: req.GetCommonAnnotations(),
		ExternalURL:       req.GetExternalURL(),
		Version:           req.GetVersion(),
		GroupKey:          req.GetGroupKey(),
		TruncatedAlerts:   req.GetTruncatedAlerts(),
	}
}

func ToAlertItems(alerts []*common.AlertItem) []*bo.AlertItem {
	if len(alerts) == 0 {
		return []*bo.AlertItem{}
	}
	return slices.MapFilter(alerts, func(alert *common.AlertItem) (*bo.AlertItem, bool) {
		if alertItem := ToAlertItem(alert); alertItem != nil {
			return alertItem, true
		}
		return nil, false
	})
}

func ToAlertItem(alert *common.AlertItem) *bo.AlertItem {
	if validate.IsNil(alert) {
		return nil
	}
	return &bo.AlertItem{
		Status:       alert.GetStatus().String(),
		Labels:       alert.GetLabels(),
		Annotations:  alert.GetAnnotations(),
		StartsAt:     alert.GetStartsAt(),
		EndsAt:       alert.GetEndsAt(),
		GeneratorURL: alert.GetGeneratorURL(),
		Fingerprint:  alert.GetFingerprint(),
		Value:        alert.GetValue(),
	}
}
