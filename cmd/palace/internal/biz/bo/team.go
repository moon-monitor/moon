package bo

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
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
}

type SaveTeamRequest interface {
	GetName() string
	GetRemark() string
	GetLogo() string
}

func NewSaveOneTeamRequest(req SaveTeamRequest) *SaveOneTeamRequest {
	return &SaveOneTeamRequest{
		name:   req.GetName(),
		remark: req.GetRemark(),
		logo:   req.GetLogo(),
	}
}

func (o *SaveOneTeamRequest) WithCreateTeamRequest(ctx context.Context) (Team, error) {
	leaderID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorPermissionDenied("user id not found in context")
	}
	o.leaderID = leaderID
	return o, nil
}

func (o *SaveOneTeamRequest) WithUpdateTeamRequest(team Team) Team {
	o.team = team
	return o
}

type SaveOneTeamRequest struct {
	team     Team
	name     string
	remark   string
	logo     string
	leaderID uint32
}

func (o *SaveOneTeamRequest) GetCreatedAt() time.Time {
	if o == nil || o.team == nil {
		return time.Now()
	}
	return o.team.GetCreatedAt()
}

func (o *SaveOneTeamRequest) GetUpdatedAt() time.Time {
	if o == nil || o.team == nil {
		return time.Now()
	}
	return o.team.GetUpdatedAt()
}

func (o *SaveOneTeamRequest) GetTeamID() uint32 {
	if o == nil || o.team == nil {
		return 0
	}
	return o.team.GetTeamID()
}

func (o *SaveOneTeamRequest) GetName() string {
	return o.name
}

func (o *SaveOneTeamRequest) GetRemark() string {
	return o.remark
}

func (o *SaveOneTeamRequest) GetLogo() string {
	return o.logo
}

func (o *SaveOneTeamRequest) GetStatus() vobj.TeamStatus {
	if o == nil || o.team == nil {
		return vobj.TeamStatusApproval
	}
	return o.team.GetStatus()
}

func (o *SaveOneTeamRequest) GetLeaderID() uint32 {
	if o == nil || o.team == nil {
		return o.leaderID
	}
	return o.team.GetLeaderID()
}

func (o *SaveOneTeamRequest) GetUUID() uuid.UUID {
	if o == nil || o.team == nil {
		return uuid.New()
	}
	return o.team.GetUUID()
}

func (o *SaveOneTeamRequest) GetCapacity() vobj.TeamCapacity {
	if o == nil || o.team == nil {
		return vobj.TeamCapacityDefault
	}
	return o.team.GetCapacity()
}

func (o *SaveOneTeamRequest) GetDBName() string {
	if o == nil || o.team == nil {
		return ""
	}
	return o.team.GetDBName()
}

type TeamListRequest struct {
	*PaginationRequest
	Keyword string
	Status  []vobj.TeamStatus
	UserIds []uint32
}

type TeamListReply struct {
	*PaginationReply
	Items []Team
}
