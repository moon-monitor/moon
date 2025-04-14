package bo

import (
	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/safety"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewSendHookParams(configs []HookConfig, opts ...SendHookParamsOption) (SendHookParams, error) {
	configs = slices.MapFilter(configs, func(configItem HookConfig) (HookConfig, bool) {
		return configItem, configItem != nil && configItem.GetEnable()
	})
	if len(configs) == 0 {
		return nil, merr.ErrorParamsError("No hook configuration is available")
	}
	params := &sendHookParams{
		Configs: configs,
	}
	for _, opt := range opts {
		if err := opt(params); err != nil {
			return nil, err
		}
	}
	return params, nil
}

type HookConfig interface {
	GetName() string
	GetApp() common.HookAPP
	GetUrl() string
	GetSecret() string
	GetToken() string
	GetUsername() string
	GetPassword() string
	GetHeaders() map[string]string
	GetEnable() bool
}

type SendHookParams interface {
	GetBody() *safety.Map[common.HookAPP, []byte]
	GetConfigs() []HookConfig
}

type sendHookParams struct {
	BodyMap *safety.Map[common.HookAPP, []byte]
	Configs []HookConfig
}

func (s *sendHookParams) GetBody() *safety.Map[common.HookAPP, []byte] {
	if s == nil {
		return nil
	}
	return s.BodyMap
}

func (s *sendHookParams) GetConfigs() []HookConfig {
	if s == nil {
		return nil
	}
	return slices.UniqueWithFunc(s.Configs, func(configItem HookConfig) string { return configItem.GetUrl() })
}

type SendHookParamsOption func(params *sendHookParams) error

func WithSendHookParamsOptionBody(bodyMap map[common.HookAPP][]byte) SendHookParamsOption {
	return func(params *sendHookParams) error {
		if bodyMap == nil {
			return merr.ErrorParamsError("body is empty").WithMetadata(map[string]string{
				"body": "body is empty",
			})
		}

		params.BodyMap = safety.NewMap(bodyMap)
		return nil
	}
}
