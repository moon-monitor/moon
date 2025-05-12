package bo

import (
	"strings"

	"github.com/moon-monitor/moon/pkg/util/validate"
)

type AlertItem struct {
	Status       string
	Labels       map[string]string
	Annotations  map[string]string
	StartsAt     string
	EndsAt       string
	GeneratorURL string
	Fingerprint  string
	Value        string
}

type AlertsItem struct {
	Receiver          string
	Status            string
	Alerts            []*AlertItem
	GroupLabels       map[string]string
	CommonLabels      map[string]string
	CommonAnnotations map[string]string
	ExternalURL       string
	Version           string
	GroupKey          string
	TruncatedAlerts   int32
}

// GetReceiver implements bo.AlertsItem.
func (a *AlertsItem) GetReceiver() []string {
	if a == nil || validate.TextIsNull(a.Receiver) {
		return []string{}
	}
	return strings.Split(a.Receiver, ",")
}
