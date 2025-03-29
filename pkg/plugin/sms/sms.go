package sms

import "context"

type Message struct {
	Content     string `json:"content"`
	Code        string `json:"code"`
	PhoneNumber string `json:"phoneNumber"`
}

type Sender interface {
	Send(ctx context.Context, message Message) error
	SendBatch(ctx context.Context, code string, messages []Message) error
}
