package impl

import (
	"context"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
)

// NewDashboardRepo creates a new dashboard repository
func NewDashboardRepo(data *data.Data, logger log.Logger) repository.Dashboard {
	return &dashboardImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.dashboard")),
	}
}

type dashboardImpl struct {
	*data.Data

	helper *log.Helper
}

// SaveDashboard save dashboard
func (r *dashboardImpl) SaveDashboard(ctx context.Context, dashboard *team.Dashboard) error {
	query, teamID, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return err
	}
	mutation := query.Dashboard
	if dashboard.ID == 0 {
		dashboard.WithContext(ctx)
		return mutation.WithContext(ctx).Create(dashboard)
	}
	wrappers := []gen.Condition{
		mutation.TeamID.Eq(teamID),
		mutation.ID.Eq(dashboard.ID),
	}
	updates := []field.AssignExpr{
		mutation.Title.Value(dashboard.Title),
		mutation.Remark.Value(dashboard.Remark),
		mutation.ColorHex.Value(dashboard.ColorHex),
	}
	_, err = mutation.WithContext(ctx).Where(wrappers...).UpdateColumnSimple(updates...)
	return err
}

// DeleteDashboard delete dashboard by id
func (r *dashboardImpl) DeleteDashboard(ctx context.Context, id uint32) error {
	query, teamID, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return err
	}
	mutation := query.Dashboard
	wrappers := []gen.Condition{
		mutation.TeamID.Eq(teamID),
		mutation.ID.Eq(id),
	}
	_, err = mutation.WithContext(ctx).Where(wrappers...).Delete()
	return err
}

// GetDashboard get dashboard by id
func (r *dashboardImpl) GetDashboard(ctx context.Context, id uint32) (*team.Dashboard, error) {
	query, teamID, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return nil, err
	}
	mutation := query.Dashboard
	wrappers := []gen.Condition{
		mutation.TeamID.Eq(teamID),
		mutation.ID.Eq(id),
	}
	dashboardDo, err := mutation.WithContext(ctx).Where(wrappers...).First()
	if err != nil {
		return nil, teamDashboardNotFound(err)
	}
	return dashboardDo, nil
}

// ListDashboards list dashboards with filter
func (r *dashboardImpl) ListDashboards(ctx context.Context, req *bo.ListDashboardReq) (*bo.ListDashboardReply, error) {
	query, teamID, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return nil, err
	}
	mutation := query.Dashboard
	wrapper := mutation.WithContext(ctx).Where(mutation.TeamID.Eq(teamID))

	if !req.Status.IsUnknown() {
		wrapper = wrapper.Where(mutation.Status.Eq(req.Status.GetValue()))
	}

	if req.PaginationRequest != nil {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		req.WithTotal(total)
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
	}

	dashboards, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	return &bo.ListDashboardReply{
		PaginationReply: req.ToReply(),
		Dashboards:      dashboards,
	}, nil
}

// BatchUpdateDashboardStatus update multiple dashboards status
func (r *dashboardImpl) BatchUpdateDashboardStatus(ctx context.Context, ids []uint32, status vobj.GlobalStatus) error {
	if len(ids) == 0 {
		return nil
	}
	query, teamID, err := getTeamBizQuery(ctx, r)
	if err != nil {
		return err
	}
	mutation := query.Dashboard
	wrappers := []gen.Condition{
		mutation.TeamID.Eq(teamID),
		mutation.ID.In(ids...),
	}
	_, err = mutation.WithContext(ctx).Where(wrappers...).UpdateColumnSimple(mutation.Status.Value(int8(status)))
	return err
}
