package do

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Dict interface {
	TeamBase
	GetKey() string
	GetValue() string
	GetStatus() vobj.GlobalStatus
	GetType() vobj.DictType
	GetColor() string
	GetLang() string
}
