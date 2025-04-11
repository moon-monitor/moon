package repository

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/bo"
)

type EventBus interface {
	InStrategyJobEventBus() chan<- bo.StrategyJob
	OutStrategyJobEventBus() <-chan bo.StrategyJob

	InAlertEventBus() chan<- bo.AlertJob
	OutAlertEventBus() <-chan bo.AlertJob
}
