package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewMemberRepo(data *data.Data, logger log.Logger) repository.Member {
	return &memberImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.member")),
	}
}

type memberImpl struct {
	*data.Data
	helper *log.Helper
}

func (m *memberImpl) List(ctx context.Context, req *bo.TeamMemberListRequest) (*bo.TeamMemberListReply, error) {
	if validate.IsNil(req) {
		return nil, merr.ErrorParamsError("invalid request")
	}
	bizQuery, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return nil, err
	}
	memberQuery := bizQuery.Member
	wrapper := memberQuery.WithContext(ctx).Where(memberQuery.TeamID.Eq(teamID))
	if !validate.TextIsNull(req.Keyword) {
		ors := []gen.Condition{
			memberQuery.MemberName.Like(req.Keyword),
			memberQuery.Remark.Like(req.Keyword),
		}
		wrapper = wrapper.Where(memberQuery.Or(ors...))
	}
	if len(req.Status) > 0 {
		status := slices.Map(req.Status, func(statusItem vobj.MemberStatus) int8 { return statusItem.GetValue() })
		wrapper = wrapper.Where(memberQuery.Status.In(status...))
	}
	if len(req.Positions) > 0 {
		positions := slices.Map(req.Positions, func(positionItem vobj.Role) int8 { return positionItem.GetValue() })
		wrapper = wrapper.Where(memberQuery.Position.In(positions...))
	}
	if validate.IsNotNil(req.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
		req.WithTotal(total)
	}
	members, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	return req.ToTeamMemberListReply(members), nil
}

func (m *memberImpl) UpdateStatus(ctx context.Context, req bo.UpdateMemberStatus) error {
	if validate.IsNil(req) {
		return merr.ErrorParamsError("invalid request")
	}
	if len(req.GetMembers()) == 0 {
		return nil
	}
	bizQuery, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return err
	}
	memberIds := slices.MapFilter(req.GetMembers(), func(member do.TeamMember) (uint32, bool) {
		if validate.IsNil(member) || member.GetID() <= 0 {
			return 0, false
		}
		return member.GetID(), true
	})
	if len(memberIds) == 0 {
		return nil
	}
	memberQuery := bizQuery.Member
	wrappers := []gen.Condition{
		memberQuery.TeamID.Eq(teamID),
		memberQuery.ID.In(memberIds...),
	}
	_, err = memberQuery.WithContext(ctx).Where(wrappers...).UpdateSimple(memberQuery.Status.Value(req.GetStatus().GetValue()))
	return err
}

func (m *memberImpl) UpdatePosition(ctx context.Context, req bo.UpdateMemberPosition) error {
	bizQuery, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return err
	}
	memberQuery := bizQuery.Member
	wrappers := []gen.Condition{
		memberQuery.TeamID.Eq(teamID),
		memberQuery.ID.Eq(req.GetMember().GetID()),
	}
	_, err = memberQuery.WithContext(ctx).Where(wrappers...).UpdateSimple(memberQuery.Position.Value(req.GetPosition().GetValue()))
	return err
}

func (m *memberImpl) UpdateRoles(ctx context.Context, req bo.UpdateMemberRoles) error {
	memberDo := &team.Member{
		TeamModel: do.TeamModel{
			CreatorModel: do.CreatorModel{
				BaseModel: do.BaseModel{ID: req.GetMember().GetID()},
			},
		},
	}

	roles := slices.MapFilter(req.GetRoles(), func(role do.TeamRole) (*team.Role, bool) {
		if validate.IsNil(role) || role.GetID() <= 0 {
			return nil, false
		}
		return &team.Role{
			TeamModel: do.TeamModel{
				CreatorModel: do.CreatorModel{
					BaseModel: do.BaseModel{ID: role.GetID()},
				},
			},
		}, true
	})
	bizMutation, _, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return err
	}
	memberMutation := bizMutation.Member
	rolesAssociation := memberMutation.Roles.WithContext(ctx).Model(memberDo)
	if len(roles) == 0 {
		return rolesAssociation.Clear()
	}
	return rolesAssociation.Replace(roles...)
}

func (m *memberImpl) Get(ctx context.Context, id uint32) (do.TeamMember, error) {
	query, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return nil, err
	}
	memberQuery := query.Member
	wrappers := []gen.Condition{
		memberQuery.TeamID.Eq(teamID),
		memberQuery.ID.Eq(id),
	}
	member, err := memberQuery.WithContext(ctx).Where(wrappers...).First()
	if err != nil {
		return nil, teamMemberNotFound(err)
	}
	return member, nil
}

func (m *memberImpl) Find(ctx context.Context, ids []uint32) ([]do.TeamMember, error) {
	if len(ids) == 0 {
		return nil, nil
	}
	query, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return nil, err
	}
	memberQuery := query.Member
	wrappers := []gen.Condition{
		memberQuery.TeamID.Eq(teamID),
		memberQuery.ID.In(ids...),
	}
	members, err := memberQuery.WithContext(ctx).Where(wrappers...).Find()
	if err != nil {
		return nil, err
	}
	memberDos := slices.Map(members, func(member *team.Member) do.TeamMember { return member })
	return memberDos, nil
}

func (m *memberImpl) FindByUserID(ctx context.Context, userID uint32) (do.TeamMember, error) {
	query, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return nil, err
	}
	memberQuery := query.Member
	wrappers := []gen.Condition{
		memberQuery.TeamID.Eq(teamID),
		memberQuery.UserID.Eq(userID),
	}
	member, err := memberQuery.WithContext(ctx).Where(wrappers...).First()
	if err != nil {
		return nil, teamMemberNotFound(err)
	}
	return member, nil
}
