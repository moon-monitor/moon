package server_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/moon-monitor/moon/pkg/plugin/server"
)

// TestNewTicker verifies that TestNewTicker correctly initializes a Ticker with the given interval and task.
func TestNewTicker(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	interval := 1 * time.Second
	start := time.Now()
	task := &server.TickTask{
		Fn: func(ctx context.Context, isStop bool) error {
			if isStop {
				t.Logf("Task stopped")
				return nil
			}
			diff := time.Now().Sub(start)
			diff = diff.Round(time.Second)
			if diff < interval {
				t.Errorf("Expected task to be executed after %v, but it was executed after %v", interval, diff)
				return fmt.Errorf("task executed after %v", diff)
			}
			t.Logf("Task executed after %v", diff)
			return nil
		},
		Name:    "定时器",
		Timeout: 0,
	}

	ticker := server.NewTicker(interval, task)
	err := ticker.Start(ctx)
	if err != nil {
		t.Fatalf("Failed to start timer: %v", err)
	}

	<-ctx.Done()
	ticker.Stop(ctx)
}

func TestTestNewTickers(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	list := []time.Duration{
		1 * time.Second,
		2 * time.Second,
		3 * time.Second,
		4 * time.Second,
		5 * time.Second,
	}
	start := time.Now()
	task := make(map[time.Duration]*server.TickTask)
	for _, v := range list {
		task[v] = &server.TickTask{
			Fn: func(ctx context.Context, isStop bool) error {
				if isStop {
					t.Logf("Task stopped")
					return nil
				}
				diff := time.Now().Sub(start)
				diff = diff.Round(time.Second)
				if diff < v {
					t.Errorf("Expected task to be executed after %v, but it was executed after %v", v, diff)
					return fmt.Errorf("task executed after %v: %v", v, diff)
				}
				t.Logf("Task executed after %v: %v", v, diff)
				return nil
			},
			Name:    fmt.Sprintf("%v", v),
			Timeout: 0,
		}
	}

	tickers := server.NewTickers(server.WithTickersTasks(task))
	err := tickers.Start(ctx)
	if err != nil {
		t.Fatalf("Failed to start timer: %v", err)
	}

	<-ctx.Done()
	tickers.Stop(ctx)
}
