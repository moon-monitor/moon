package do

import (
	"github.com/google/uuid"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/config"
)

type Team interface {
	Creator
	GetName() string
	GetRemark() string
	GetLogo() string
	GetStatus() vobj.TeamStatus
	GetLeaderID() uint32
	GetUUID() uuid.UUID
	GetCapacity() vobj.TeamCapacity
	GetBizDBConfig() *config.Database
	GetAlarmDBConfig() *config.Database
	GetLeader() User
	GetAdmins() []User
	GetResources() []Resource
}
