package ali

import (
	"github.com/go-kratos/kratos/v2/log"
)

func WithAliyunLogger(logger log.Logger) AliyunOption {
	return func(a *aliyun) {
		a.helper = log.NewHelper(log.With(logger, "module", "plugin.sms.aliyun"))
	}
}

func WithAliyunSignName(signName string) AliyunOption {
	return func(a *aliyun) {
		a.signName = signName
	}
}

func WithAliyunEndpoint(endpoint string) AliyunOption {
	return func(a *aliyun) {
		a.endpoint = endpoint
	}
}
