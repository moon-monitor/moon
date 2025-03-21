package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
)

var _ transport.Server = (*Ticker)(nil)

func NewTicker(interval time.Duration, task TickTask, opts ...TickerOption) *Ticker {
	t := &Ticker{
		interval: interval,
		stop:     make(chan struct{}),
		task:     task,
	}
	for _, opt := range opts {
		opt(t)
	}
	if t.helper == nil {
		WithTickerLogger(log.DefaultLogger)(t)
	}
	return t
}

type TickTask func() error

type Ticker struct {
	interval time.Duration
	ticker   *time.Ticker
	stop     chan struct{}
	task     TickTask

	helper *log.Helper
}

type TickerOption func(*Ticker)

func WithTickerLogger(logger log.Logger) TickerOption {
	return func(t *Ticker) {
		t.helper = log.NewHelper(log.With(logger, "module", "server.tick"))
	}
}

func (t *Ticker) Start(ctx context.Context) error {
	t.ticker = time.NewTicker(t.interval)
	go func() {
		for {
			select {
			case <-t.ticker.C:
				t.task()
			case <-t.stop:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (t *Ticker) Stop(ctx context.Context) error {
	close(t.stop)
	t.ticker.Stop()
	return nil
}

var _ transport.Server = (*Tickers)(nil)

func NewTickers(tasks map[time.Duration]TickTask) *Tickers {
	return &Tickers{
		tasks:   tasks,
		tickers: make([]*Ticker, 0, len(tasks)),
	}
}

type Tickers struct {
	tasks   map[time.Duration]TickTask
	tickers []*Ticker
}

func (t *Tickers) Start(ctx context.Context) error {
	for interval, task := range t.tasks {
		t.tickers = append(t.tickers, NewTicker(interval, task))
	}
	for _, ticker := range t.tickers {
		ticker.Start(ctx)
	}
	return nil
}

func (t *Tickers) Stop(ctx context.Context) error {
	for _, ticker := range t.tickers {
		ticker.Stop(ctx)
	}
	return nil
}
