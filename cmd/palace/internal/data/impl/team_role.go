package impl

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"
	
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewTeamRole(data *data.Data) repository.TeamRole {
	return &teamRoleImpl{
		Data: data,
	}
}

type teamRoleImpl struct {
	*data.Data
}

func (t *teamRoleImpl) Get(ctx context.Context, id uint32) (do.TeamRole, error) {
	bizQuery, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}
	roleQuery := bizQuery.Role
	wrapper := []gen.Condition{
		roleQuery.TeamID.Eq(teamID),
		roleQuery.ID.Eq(id),
	}
	return roleQuery.WithContext(ctx).Where(wrapper...).First()
}

func (t *teamRoleImpl) List(ctx context.Context, req *bo.ListRoleReq) (*bo.ListRoleReply, error) {
	bizQuery, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return nil, err
	}
	roleQuery := bizQuery.Role
	wrapper := roleQuery.WithContext(ctx).Where(roleQuery.TeamID.Eq(teamID))

	if !req.Status.IsUnknown() {
		wrapper = wrapper.Where(roleQuery.Status.Eq(req.Status.GetValue()))
	}
	if !validate.TextIsNull(req.Keyword) {
		wrapper = wrapper.Where(roleQuery.Name.Like(req.Keyword))
	}
	if validate.IsNotNil(req.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
		req.WithTotal(total)
	}
	roles, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	return req.ToListTeamRoleReply(roles), nil
}

func (t *teamRoleImpl) Create(ctx context.Context, role bo.Role) error {
	teamDo := &team.Role{
		Name:   role.GetName(),
		Remark: role.GetRemark(),
		Status: role.GetStatus(),
		Menus: slices.MapFilter(role.GetMenus(), func(menu do.Menu) (*team.Menu, bool) {
			if validate.IsNil(menu) || menu.GetID() <= 0 {
				return nil, false
			}
			return &team.Menu{
				TeamModel: do.TeamModel{
					CreatorModel: do.CreatorModel{
						BaseModel: do.BaseModel{ID: menu.GetID()},
					},
				},
			}, true
		}),
	}
	teamDo.WithContext(ctx)
	bizQuery, _, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	bizRoleQuery := bizQuery.Role
	return bizRoleQuery.WithContext(ctx).Create(teamDo)
}

func (t *teamRoleImpl) Update(ctx context.Context, role bo.Role) error {
	bizQuery, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	bizRoleQuery := bizQuery.Role
	wrapper := []gen.Condition{
		bizRoleQuery.TeamID.Eq(teamID),
		bizRoleQuery.ID.Eq(role.GetID()),
	}
	mutations := []field.AssignExpr{
		bizRoleQuery.Name.Value(role.GetName()),
		bizRoleQuery.Remark.Value(role.GetRemark()),
		bizRoleQuery.Status.Value(role.GetStatus().GetValue()),
	}
	_, err = bizRoleQuery.WithContext(ctx).Where(wrapper...).UpdateColumnSimple(mutations...)
	if err != nil {
		return err
	}

	roleDo := &team.Role{
		TeamModel: do.TeamModel{
			CreatorModel: do.CreatorModel{
				BaseModel: do.BaseModel{ID: role.GetID()},
			},
		},
	}
	menuDos := slices.MapFilter(role.GetMenus(), func(menu do.Menu) (*team.Menu, bool) {
		if validate.IsNil(menu) || menu.GetID() <= 0 {
			return nil, false
		}
		return &team.Menu{
			TeamModel: do.TeamModel{
				CreatorModel: do.CreatorModel{
					BaseModel: do.BaseModel{ID: menu.GetID()},
				},
			},
		}, true
	})
	menuMutation := bizRoleQuery.Menus.WithContext(ctx).Model(roleDo)
	if len(menuDos) == 0 {
		return menuMutation.Clear()
	}
	return menuMutation.Replace(menuDos...)
}

func (t *teamRoleImpl) Delete(ctx context.Context, id uint32) error {
	bizQuery, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	bizRoleQuery := bizQuery.Role
	wrapper := []gen.Condition{
		bizRoleQuery.TeamID.Eq(teamID),
		bizRoleQuery.ID.Eq(id),
	}
	_, err = bizRoleQuery.WithContext(ctx).Where(wrapper...).Delete()
	return err
}

func (t *teamRoleImpl) UpdateStatus(ctx context.Context, req *bo.UpdateTeamRoleStatusReq) error {
	bizQuery, teamID, err := getTeamBizQuery(ctx, t)
	if err != nil {
		return err
	}
	bizRoleQuery := bizQuery.Role
	wrapper := []gen.Condition{
		bizRoleQuery.TeamID.Eq(teamID),
		bizRoleQuery.ID.Eq(req.RoleID),
	}
	_, err = bizRoleQuery.WithContext(ctx).Where(wrapper...).UpdateColumnSimple(bizRoleQuery.Status.Value(req.Status.GetValue()))
	return err
}
