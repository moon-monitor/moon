package sms

import (
	"context"
	"fmt"
	"plugin"

	"github.com/go-kratos/kratos/v2/log"
)

type Message struct {
	TemplateParam string `json:"templateParam"`
	TemplateCode  string `json:"templateCode"`
}

type Sender interface {
	Send(ctx context.Context, phoneNumber string, message Message) error
	SendBatch(ctx context.Context, phoneNumbers []string, messages Message) error
}

// PluginNewFunc is the function signature plugins must implement
type PluginNewFunc func(log.Logger) (Sender, error)

func LoadPlugin(path string, logger log.Logger) (Sender, error) {
	p, err := plugin.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not load plugin: %v", err)
	}

	newFuncSym, err := p.Lookup("New")
	if err != nil {
		return nil, fmt.Errorf("could not find New symbol: %v", err)
	}

	newFunc, ok := newFuncSym.(func(log.Logger) (Sender, error))
	if !ok {
		return nil, fmt.Errorf("plugin New has wrong signature")
	}

	return newFunc(logger)
}
