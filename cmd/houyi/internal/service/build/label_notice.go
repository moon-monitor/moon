package build

import (
	"github.com/moon-monitor/moon/cmd/houyi/internal/biz/do"
)

type LabelNotice interface {
	GetKey() string
	GetValue() string
	GetReceiverRoutes() []string
}

func ToLabelNotice(notice LabelNotice) *do.LabelNotices {
	return &do.LabelNotices{
		Key:            notice.GetKey(),
		Value:          notice.GetValue(),
		ReceiverRoutes: notice.GetReceiverRoutes(),
	}
}
