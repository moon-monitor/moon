package do

import (
	"time"

	"github.com/google/uuid"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
)

type Team interface {
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetTeamID() uint32
	GetName() string
	GetRemark() string
	GetLogo() string
	GetStatus() vobj.TeamStatus
	GetLeaderID() uint32
	GetUUID() uuid.UUID
	GetCapacity() vobj.TeamCapacity
	GetDBName() string
	GetLeader() User
	GetAdmins() []User
	GetResources() []Resource
}
