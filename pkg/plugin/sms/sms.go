package sms

import "context"

type Sender interface {
	Send(ctx context.Context, message Message) error
	SendBatch(ctx context.Context, code string, messages []Message) error
}

type Message struct {
	Content     string
	Code        string
	PhoneNumber string
}
