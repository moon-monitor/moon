package server

import (
	"context"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
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
				t.call(ctx)
			case <-t.stop:
				return
			case <-ctx.Done():
				return
			}
		}
	}()
	return nil
}

func (t *Ticker) call(ctx context.Context) {
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

func NewTickers(opts ...TickersOption) *Tickers {
	t := &Tickers{
		autoID:  uint64(1),
		tickers: make(map[uint64]*Ticker),
		recycle: make([]uint64, 0, 100),
		logger:  log.DefaultLogger,
	}
	for _, opt := range opts {
		opt(t)
	}

	return t
}

type Tickers struct {
	mu      sync.RWMutex
	autoID  uint64
	recycle []uint64
	tickers map[uint64]*Ticker
	logger  log.Logger
}

type TickersOption func(*Tickers)

func WithTickersLogger(logger log.Logger) TickersOption {
	return func(t *Tickers) {
		t.logger = logger
	}
}

func WithTickersTasks(tasks map[time.Duration]*TickTask) TickersOption {
	return func(t *Tickers) {
		for interval, task := range tasks {
			t.Add(interval, task)
		}
	}
}

func (t *Tickers) Add(interval time.Duration, task *TickTask) uint64 {
	t.mu.Lock()
	defer t.mu.Unlock()
	id := t.autoID
	if len(t.recycle) > 0 {
		id = t.recycle[0]
		t.recycle = t.recycle[1:]
	} else {
		t.autoID++
	}
	ticker := NewTicker(interval, task, WithTickerLogger(t.logger))
	defer ticker.Start(context.Background())
	t.tickers[id] = ticker
	return id
}

func (t *Tickers) Remove(id uint64) {
	t.mu.Lock()
	defer t.mu.Unlock()
	ticker, ok := t.tickers[id]
	if !ok {
		return
	}
	ticker.Stop(context.Background())
	delete(t.tickers, id)
	t.recycle = append(t.recycle, id)
}

func (t *Tickers) Start(ctx context.Context) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, ticker := range t.tickers {
		ticker.Start(ctx)
	}
	return nil
}

func (t *Tickers) Stop(ctx context.Context) error {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, ticker := range t.tickers {
		ticker.Stop(ctx)
	}
	return nil
}
