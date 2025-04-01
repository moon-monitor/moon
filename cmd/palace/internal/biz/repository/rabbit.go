package repository

import (
	rabbitv1 "github.com/moon-monitor/moon/pkg/api/rabbit/v1"
)

type Rabbit interface {
	Send() (rabbitv1.SendClient, bool)
	Sync() (rabbitv1.SyncClient, bool)
}
