package repository

type EventBus interface {
	InMetricIDEventBus() chan<- string
	OutMetricIDEventBus() <-chan string
}
