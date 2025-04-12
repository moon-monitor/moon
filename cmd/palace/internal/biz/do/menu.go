package do

import (
	"time"

	"gorm.io/plugin/soft_delete"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Menu interface {
	GetID() uint32
	GetName() string
	GetPath() string
	GetStatus() vobj.GlobalStatus
	GetIcon() string
	GetParentID() uint32
	GetType() vobj.MenuType
	GetResources() []Resource
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetDeletedAt() soft_delete.DeletedAt
	GetParent() Menu
}
