package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewResourceRepo(d *data.Data, logger log.Logger) repository.Resource {
	return &resourceImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.resource")),
	}
}

type resourceImpl struct {
	*data.Data
	helper *log.Helper
}

func (r *resourceImpl) GetResources(ctx context.Context) ([]do.Resource, error) {
	mainQuery := getMainQuery(ctx, r)
	resources, err := mainQuery.Resource.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return slices.Map(resources, func(resource *system.Resource) do.Resource { return resource }), nil
}

func (r *resourceImpl) GetResourceByID(ctx context.Context, id uint32) (do.Resource, error) {
	mainQuery := getMainQuery(ctx, r)
	resource := mainQuery.Resource
	resourceDo, err := resource.WithContext(ctx).Where(resource.ID.Eq(id)).First()
	if err != nil {
		return nil, resourceNotFound(err)
	}
	return resourceDo, nil
}

func (r *resourceImpl) GetResourceByOperation(ctx context.Context, operation string) (do.Resource, error) {
	mainQuery := getMainQuery(ctx, r)
	resource := mainQuery.Resource
	resourceDo, err := resource.WithContext(ctx).Where(resource.Path.Eq(operation)).First()
	if err != nil {
		return nil, resourceNotFound(err)
	}
	return resourceDo, nil
}

func (r *resourceImpl) BatchUpdateResourceStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error {
	if len(ids) == 0 {
		return nil
	}
	mainQuery := getMainQuery(ctx, r)
	resource := mainQuery.Resource
	_, err := resource.WithContext(ctx).
		Where(resource.ID.In(ids...)).
		Update(resource.Status, int8(status))
	return err
}

func (r *resourceImpl) ListResources(ctx context.Context, req *bo.ListResourceReq) (*bo.ListResourceReply, error) {
	mainQuery := getMainQuery(ctx, r)
	resource := mainQuery.Resource
	resourceQuery := resource.WithContext(ctx)
	if len(req.Statuses) > 0 {
		resourceQuery = resourceQuery.Where(resource.Status.In(slices.Map(req.Statuses, func(status vobj.GlobalStatus) int8 { return status.GetValue() })...))
	}
	if !validate.TextIsNull(req.Keyword) {
		resourceQuery = resourceQuery.Where(resource.Name.Like(req.Keyword))
	}
	if req.PaginationRequest != nil {
		total, err := resourceQuery.Count()
		if err != nil {
			return nil, err
		}
		req.WithTotal(total)
		resourceQuery = resourceQuery.Offset(req.Offset()).Limit(int(req.Limit))
	}
	resources, err := resourceQuery.Find()
	if err != nil {
		return nil, err
	}
	return req.ToListResourceReply(resources), nil
}

func (r *resourceImpl) GetMenusByUserID(ctx context.Context, userID uint32) ([]do.Menu, error) {
	mainQuery := getMainQuery(ctx, r)
	user := mainQuery.User
	userQuery := user.WithContext(ctx).Where(user.ID.Eq(userID)).Preload(user.Roles.Menus)
	userDo, err := userQuery.First()
	if err != nil {
		return nil, userNotFound(err)
	}
	menus := make([]do.Menu, 0, len(userDo.Roles))
	for _, role := range userDo.Roles {
		menus = append(menus, slices.Map(role.Menus, func(menu *system.Menu) do.Menu { return menu })...)
	}
	teamBizQuery, teamID, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return nil, err
	}
	teamMember := teamBizQuery.Member
	teamQuery := teamMember.WithContext(ctx).Where(teamMember.TeamID.Eq(teamID)).Preload(teamMember.Roles.Menus)
	teamMemberDo, err := teamQuery.First()
	if err != nil {
		return nil, teamMemberNotFound(err)
	}
	for _, role := range teamMemberDo.Roles {
		menus = append(menus, slices.Map(role.Menus, func(menu *team.Menu) do.Menu { return menu })...)
	}
	return menus, nil
}

func (r *resourceImpl) GetMenus(ctx context.Context, t vobj.MenuType) ([]do.Menu, error) {
	mainQuery := getMainQuery(ctx, r)
	menu := mainQuery.Menu
	menus, err := menu.WithContext(ctx).Where(menu.Type.Eq(t.GetValue())).Find()
	if err != nil {
		return nil, err
	}
	return slices.Map(menus, func(menu *system.Menu) do.Menu { return menu }), nil
}
