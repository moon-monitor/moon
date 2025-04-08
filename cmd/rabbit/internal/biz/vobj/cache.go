package vobj

import (
	"github.com/moon-monitor/moon/pkg/plugin/cache"
)

const (
	EmailCacheKey cache.K = "rabbit:config:email"
	SmsCacheKey   cache.K = "rabbit:config:sms"
	HookCacheKey  cache.K = "rabbit:config:hook"
)
