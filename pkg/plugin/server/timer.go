package server

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"

	"github.com/moon-monitor/moon/pkg/util/safety"
)

var _ transport.Server = (*Ticker)(nil)

func NewTicker(interval time.Duration, task *TickTask, opts ...TickerOption) *Ticker {
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

type TickTask struct {
	Fn      func(ctx context.Context) error
	Name    string
	Timeout time.Duration
}

type Ticker struct {
	interval time.Duration
	ticker   *time.Ticker
	stop     chan struct{}
	task     *TickTask

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
				t.Call(ctx)
			case <-t.stop:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (t *Ticker) Call(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, t.task.Timeout)
	defer cancel()
	if err := t.task.Fn(ctx); err != nil {
		t.helper.Errorf("execute task %s error: %v", t.task.Name, err)
	}
}

func (t *Ticker) Stop(ctx context.Context) error {
	close(t.stop)
	t.ticker.Stop()
	return nil
}

var _ transport.Server = (*Tickers)(nil)

func NewTickers(tasks map[time.Duration]*TickTask) *Tickers {
	id := safety.NewInt64(0)
	tickerMap := safety.NewMap[int64, *Ticker]()
	recycle := safety.NewSlice[int64](100)
	for interval, task := range tasks {
		tickerMap.Set(id.Inc(), NewTicker(interval, task))
	}
	return &Tickers{
		autoID:  id,
		tickers: tickerMap,
		recycle: recycle,
	}
}

type Tickers struct {
	autoID  *safety.Int64
	recycle *safety.Slice[int64]
	tickers *safety.Map[int64, *Ticker]
}

func (t *Tickers) Add(interval time.Duration, task *TickTask) int64 {
	id, ok := t.recycle.Pop()
	if !ok {
		id = t.autoID.Inc()
	}
	ticker := NewTicker(interval, task)
	defer ticker.Start(context.Background())
	t.tickers.Set(id, ticker)
	return id
}

func (t *Tickers) Remove(id int64) {
	ticker, ok := t.tickers.Get(id)
	if !ok {
		return
	}
	ticker.Stop(context.Background())
	t.tickers.Delete(id)
	t.recycle.Append(id)
}

func (t *Tickers) Start(ctx context.Context) error {
	for _, ticker := range t.tickers.List() {
		ticker.Start(ctx)
	}
	return nil
}

func (t *Tickers) Stop(ctx context.Context) error {
	for _, ticker := range t.tickers.List() {
		ticker.Stop(ctx)
	}
	return nil
}
