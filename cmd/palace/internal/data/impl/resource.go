package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/util/slices"
)

func NewResourceRepo(d *data.Data, logger log.Logger) repository.Resource {
	return &resourceImpl{
		Data:   d,
		Query:  systemgen.Use(d.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.resource")),
	}
}

type resourceImpl struct {
	*data.Data
	*systemgen.Query
	helper *log.Helper
}

func (r *resourceImpl) GetResources(ctx context.Context) ([]*system.Resource, error) {
	return r.Resource.WithContext(ctx).Find()
}

func (r *resourceImpl) GetResourceByID(ctx context.Context, id uint32) (*system.Resource, error) {
	resourceDo, err := r.Resource.WithContext(ctx).Where(r.Resource.ID.Eq(id)).First()
	if err != nil {
		return nil, resourceNotFound(err)
	}
	return resourceDo, nil
}

func (r *resourceImpl) GetResourceByOperation(ctx context.Context, operation string) (*system.Resource, error) {
	resourceDo, err := r.Resource.WithContext(ctx).Where(r.Resource.Path.Eq(operation)).First()
	if err != nil {
		return nil, resourceNotFound(err)
	}
	return resourceDo, nil
}

func (r *resourceImpl) BatchUpdateResourceStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error {
	if len(ids) == 0 {
		return nil
	}
	_, err := r.Resource.WithContext(ctx).
		Where(r.Resource.ID.In(ids...)).
		Update(r.Resource.Status, int8(status))
	return err
}

func (r *resourceImpl) ListResources(ctx context.Context, req *bo.ListResourceReq) (*bo.ListResourceReply, error) {
	resource := r.Resource
	resourceQuery := resource.WithContext(ctx)
	if len(req.Statuses) > 0 {
		resourceQuery = resourceQuery.Where(resource.Status.In(slices.Map(req.Statuses, func(status vobj.GlobalStatus) int8 { return status.GetValue() })...))
	}
	if req.Keyword != "" {
		keyword := "%" + req.Keyword + "%"
		resourceQuery = resourceQuery.Where(resource.Name.Like(keyword))
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
	return &bo.ListResourceReply{
		PaginationReply: req.ToReply(),
		Resources:       resources,
	}, nil
}

func (r *resourceImpl) GetMenusByUserID(ctx context.Context, userID uint32) ([]*system.Menu, error) {
	user := r.User
	userQuery := user.WithContext(ctx).Where(user.ID.Eq(userID)).Preload(user.Roles.Menus)
	userDo, err := userQuery.First()
	if err != nil {
		return nil, userNotFound(err)
	}
	menus := make([]*system.Menu, 0, len(userDo.Roles))
	for _, role := range userDo.Roles {
		menus = append(menus, role.Menus...)
	}
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok || teamID <= 0 {
		return menus, nil
	}
	teamMember := r.TeamMember
	teamQuery := teamMember.WithContext(ctx).Where(teamMember.TeamID.Eq(teamID)).Preload(teamMember.Roles.Menus)
	teamMemberDo, err := teamQuery.First()
	if err != nil {
		return nil, teamMemberNotFound(err)
	}
	for _, role := range teamMemberDo.Roles {
		menus = append(menus, role.Menus...)
	}
	return menus, nil
}

func (r *resourceImpl) GetMenus(ctx context.Context, t vobj.MenuType) ([]*system.Menu, error) {
	return r.Menu.WithContext(ctx).Where(r.Menu.Type.Eq(t.GetValue())).Find()
}
