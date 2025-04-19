package do

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type OperateLog interface {
	Creator
	GetOperateType() vobj.OperateType
	GetOperateModule() vobj.ResourceModule
	GetOperateDataID() uint32
	GetOperateDataName() string
	GetTitle() string
	GetBefore() string
	GetAfter() string
	GetIP() string
}
