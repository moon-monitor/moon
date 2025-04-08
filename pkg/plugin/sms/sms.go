package sms

import "context"

type Message struct {
	Content string `json:"content"`
	Code    string `json:"code"`
}

type Sender interface {
	Send(ctx context.Context, phoneNumber string, message Message) error
	SendBatch(ctx context.Context, phoneNumbers []string, messages Message) error
}
