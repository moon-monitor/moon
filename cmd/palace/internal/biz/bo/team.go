package bo

import (
	"context"

	"github.com/google/uuid"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
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

type TeamMemberListRequest struct {
	*PaginationRequest
	Keyword   string
	Status    []vobj.MemberStatus
	Positions []vobj.Role
	TeamId    uint32
}

func (r *TeamMemberListRequest) ToTeamMemberListReply(items []*team.Member) *TeamMemberListReply {
	return &TeamMemberListReply{
		PaginationReply: r.ToReply(),
		Items:           slices.Map(items, func(member *team.Member) do.TeamMember { return member }),
	}
}

type TeamMemberListReply = ListReply[do.TeamMember]

type UpdateMemberPosition interface {
	GetMember() do.TeamMember
	GetPosition() vobj.Role
}

type UpdateMemberPositionReq struct {
	operator do.TeamMember
	member   do.TeamMember
	MemberID uint32
	Position vobj.Role
}

func (r *UpdateMemberPositionReq) GetMember() do.TeamMember {
	if r == nil {
		return nil
	}
	return r.member
}

func (r *UpdateMemberPositionReq) GetPosition() vobj.Role {
	if r == nil {
		return vobj.RoleUnknown
	}
	return r.Position
}

func (r *UpdateMemberPositionReq) WithOperator(operator do.TeamMember) *UpdateMemberPositionReq {
	r.operator = operator
	return r
}

func (r *UpdateMemberPositionReq) WithMember(member do.TeamMember) *UpdateMemberPositionReq {
	r.member = member
	return r
}

func (r *UpdateMemberPositionReq) Validate() error {
	if validate.IsNil(r.operator) {
		return merr.ErrorParamsError("invalid operator")
	}
	if validate.IsNil(r.member) {
		return merr.ErrorParamsError("invalid member")
	}
	if r.Position.IsUnknown() {
		return merr.ErrorParamsError("invalid position")
	}
	operatorPosition := r.operator.GetPosition()
	if !(operatorPosition.GT(r.Position) && operatorPosition.IsAdminOrSuperAdmin()) {
		return merr.ErrorParamsError("invalid position")
	}
	return nil
}

type UpdateMemberStatus interface {
	GetMembers() []do.TeamMember
	GetStatus() vobj.MemberStatus
}

type UpdateMemberStatusReq struct {
	operator  do.TeamMember
	members   []do.TeamMember
	MemberIds []uint32
	Status    vobj.MemberStatus
}

func (r *UpdateMemberStatusReq) GetMembers() []do.TeamMember {
	if r == nil {
		return nil
	}
	return r.members
}

func (r *UpdateMemberStatusReq) GetStatus() vobj.MemberStatus {
	if r == nil {
		return vobj.MemberStatusUnknown
	}
	return r.Status
}

func (r *UpdateMemberStatusReq) WithOperator(operator do.TeamMember) *UpdateMemberStatusReq {
	r.operator = operator
	return r
}

func (r *UpdateMemberStatusReq) WithMembers(members []do.TeamMember) *UpdateMemberStatusReq {
	r.members = slices.MapFilter(members, func(member do.TeamMember) (do.TeamMember, bool) {
		if validate.IsNil(member) || member.GetID() <= 0 {
			return nil, false
		}
		return member, true
	})
	return r
}

func (r *UpdateMemberStatusReq) Validate() error {
	if validate.IsNil(r.operator) {
		return merr.ErrorParamsError("invalid operator")
	}
	if len(r.members) == 0 {
		return merr.ErrorParamsError("invalid members")
	}
	if r.Status.IsUnknown() {
		return merr.ErrorParamsError("invalid status")
	}
	operatorPosition := r.operator.GetPosition()
	for _, member := range r.members {
		if !(operatorPosition.GT(member.GetPosition()) && operatorPosition.IsAdminOrSuperAdmin()) {
			return merr.ErrorParamsError("invalid position")
		}
	}
	return nil
}

type UpdateMemberRoles interface {
	GetMember() do.TeamMember
	GetRoles() []do.TeamRole
}

type UpdateMemberRolesReq struct {
	operator do.TeamMember
	member   do.TeamMember
	roles    []do.TeamRole
	MemberId uint32
	RoleIds  []uint32
}

func (r *UpdateMemberRolesReq) GetMember() do.TeamMember {
	if r == nil {
		return nil
	}
	return r.member
}

func (r *UpdateMemberRolesReq) GetRoles() []do.TeamRole {
	if r == nil {
		return nil
	}
	return r.roles
}

func (r *UpdateMemberRolesReq) WithOperator(operator do.TeamMember) *UpdateMemberRolesReq {
	r.operator = operator
	return r
}

func (r *UpdateMemberRolesReq) WithMember(member do.TeamMember) *UpdateMemberRolesReq {
	r.member = member
	return r
}

func (r *UpdateMemberRolesReq) WithRoles(roles []do.TeamRole) *UpdateMemberRolesReq {
	r.roles = slices.MapFilter(roles, func(role do.TeamRole) (do.TeamRole, bool) {
		if validate.IsNil(role) || role.GetID() <= 0 {
			return nil, false
		}
		return role, true
	})
	return r
}

func (r *UpdateMemberRolesReq) Validate() error {
	if validate.IsNil(r.operator) {
		return merr.ErrorParamsError("invalid operator")
	}
	if validate.IsNil(r.member) {
		return merr.ErrorParamsError("invalid member")
	}
	operatorPosition := r.operator.GetPosition()
	if !(operatorPosition.GT(r.member.GetPosition()) && operatorPosition.IsAdminOrSuperAdmin()) {
		return merr.ErrorParamsError("invalid position")
	}
	return nil
}

type RemoveMemberReq struct {
	operator do.TeamMember
	member   do.TeamMember
	MemberID uint32
}

func (r *RemoveMemberReq) GetMembers() []do.TeamMember {
	if r == nil {
		return nil
	}
	return []do.TeamMember{r.member}
}

func (r *RemoveMemberReq) GetStatus() vobj.MemberStatus {
	if r == nil {
		return vobj.MemberStatusUnknown
	}
	return vobj.MemberStatusDeleted
}

func (r *RemoveMemberReq) WithOperator(operator do.TeamMember) *RemoveMemberReq {
	r.operator = operator
	return r
}

func (r *RemoveMemberReq) WithMember(member do.TeamMember) *RemoveMemberReq {
	r.member = member
	return r
}

func (r *RemoveMemberReq) Validate() error {
	if validate.IsNil(r.operator) {
		return merr.ErrorParamsError("invalid operator")
	}
	if validate.IsNil(r.member) {
		return merr.ErrorParamsError("invalid member")
	}
	operatorPosition := r.operator.GetPosition()
	if !(operatorPosition.GT(r.member.GetPosition()) && operatorPosition.IsAdminOrSuperAdmin()) {
		return merr.ErrorParamsError("invalid position")
	}
	return nil
}
