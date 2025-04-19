package bo

import (
	"context"

	"github.com/google/uuid"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/pkg/util/slices"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

type SaveTeamRequest interface {
	GetName() string
	GetRemark() string
	GetLogo() string
}

func NewSaveOneTeamRequest(req SaveTeamRequest, id ...uint32) *SaveOneTeamRequest {
	s := &SaveOneTeamRequest{
		name:     req.GetName(),
		remark:   req.GetRemark(),
		logo:     req.GetLogo(),
		leaderID: 0,
	}
	if len(id) > 0 {
		s.id = id[0]
	}
	return s
}

func (o *SaveOneTeamRequest) WithCreateTeamRequest(ctx context.Context) (do.Team, error) {
	leaderID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorPermissionDenied("user id not found in context")
	}
	o.leaderID = leaderID
	return o, nil
}

func (o *SaveOneTeamRequest) WithUpdateTeamRequest(team do.Team) do.Team {
	o.Team = team
	return o
}

type SaveOneTeamRequest struct {
	do.Team
	id       uint32
	name     string
	remark   string
	logo     string
	leaderID uint32
}

func (o *SaveOneTeamRequest) GetID() uint32 {
	if o == nil || o.Team == nil {
		return o.id
	}
	return o.Team.GetID()
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
	if o == nil || o.Team == nil {
		return vobj.TeamStatusApproval
	}
	return o.Team.GetStatus()
}

func (o *SaveOneTeamRequest) GetLeaderID() uint32 {
	if o == nil || o.Team == nil {
		return o.leaderID
	}
	return o.Team.GetLeaderID()
}

func (o *SaveOneTeamRequest) GetUUID() uuid.UUID {
	if o == nil || o.Team == nil {
		return uuid.New()
	}
	return o.Team.GetUUID()
}

func (o *SaveOneTeamRequest) GetCapacity() vobj.TeamCapacity {
	if o == nil || o.Team == nil {
		return vobj.TeamCapacityDefault
	}
	return o.Team.GetCapacity()
}

type TeamListRequest struct {
	*PaginationRequest
	Keyword   string
	Status    []vobj.TeamStatus
	UserIds   []uint32
	LeaderId  uint32
	CreatorId uint32
}

func (r *TeamListRequest) ToTeamListReply(items []*system.Team) *TeamListReply {
	return &TeamListReply{
		PaginationReply: r.ToReply(),
		Items:           slices.Map(items, func(team *system.Team) do.Team { return team }),
	}
}

type TeamListReply = ListReply[do.Team]
