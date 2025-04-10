package repository

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/event"
)

type EventBus interface {
	InStrategyJobEventBus() chan<- event.StrategyJob
	OutStrategyJobEventBus() <-chan event.StrategyJob
}
