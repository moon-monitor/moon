package do

import (
	"encoding/json"

	"github.com/moon-monitor/moon/pkg/api/rabbit/common"
	"github.com/moon-monitor/moon/pkg/plugin/cache"
)

var _ cache.Object = (*HookConfig)(nil)

type HookConfig struct {
	Name     string            `json:"name"`
	App      common.HookAPP    `json:"app"`
	Url      string            `json:"url"`
	Secret   string            `json:"secret"`
	Token    string            `json:"token"`
	Username string            `json:"username"`
	Password string            `json:"password"`
	Headers  map[string]string `json:"headers"`
	Enable   bool              `json:"enable"`
}

func (h *HookConfig) GetName() string {
	if h == nil {
		return ""
	}
	return h.Name
}

func (h *HookConfig) GetApp() common.HookAPP {
	if h == nil {
		return 0
	}
	return h.App
}

func (h *HookConfig) GetUrl() string {
	if h == nil {
		return ""
	}
	return h.Url
}

func (h *HookConfig) GetSecret() string {
	if h == nil {
		return ""
	}
	return h.Secret
}

func (h *HookConfig) GetToken() string {
	if h == nil {
		return ""
	}
	return h.Token
}

func (h *HookConfig) GetUsername() string {
	if h == nil {
		return ""
	}
	return h.Username
}

func (h *HookConfig) GetPassword() string {
	if h == nil {
		return ""
	}
	return h.Password
}

func (h *HookConfig) GetHeaders() map[string]string {
	if h == nil {
		return nil
	}
	return h.Headers
}

func (h *HookConfig) GetEnable() bool {
	if h == nil {
		return false
	}
	return h.Enable
}

func (h *HookConfig) MarshalBinary() (data []byte, err error) {
	return json.Marshal(h)
}

func (h *HookConfig) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, h)
}

func (h *HookConfig) UniqueKey() string {
	return h.Name
}
