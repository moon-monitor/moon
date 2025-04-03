package impl

import (
	"context"

	"gorm.io/gen/field"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/teamgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
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

func (r *dashboardImpl) getDashboardTX(ctx context.Context) (*teamgen.Query, error) {
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorInternalServerError("team id not found")
	}
	db, err := r.GetBizDB(teamID)
	if err != nil {
		return nil, err
	}
	return teamgen.Use(db.GetDB()), nil
}

// SaveDashboard save dashboard
func (r *dashboardImpl) SaveDashboard(ctx context.Context, dashboard *team.Dashboard) error {
	tx, err := r.getDashboardTX(ctx)
	if err != nil {
		return err
	}
	mutation := tx.Dashboard
	if dashboard.ID == 0 {
		return mutation.WithContext(ctx).Create(dashboard)
	}
	updates := []field.AssignExpr{
		mutation.Title.Value(dashboard.Title),
		mutation.Remark.Value(dashboard.Remark),
		mutation.ColorHex.Value(dashboard.ColorHex),
	}
	_, err = mutation.WithContext(ctx).Where(mutation.ID.Eq(dashboard.ID)).UpdateColumnSimple(updates...)
	return err
}

// DeleteDashboard delete dashboard by id
func (r *dashboardImpl) DeleteDashboard(ctx context.Context, id uint32) error {
	tx, err := r.getDashboardTX(ctx)
	if err != nil {
		return err
	}
	mutation := tx.Dashboard
	_, err = mutation.WithContext(ctx).Where(mutation.ID.Eq(id)).Delete()
	return err
}

// GetDashboard get dashboard by id
func (r *dashboardImpl) GetDashboard(ctx context.Context, id uint32) (*team.Dashboard, error) {
	tx, err := r.getDashboardTX(ctx)
	if err != nil {
		return nil, err
	}
	mutation := tx.Dashboard
	dashboardDo, err := mutation.WithContext(ctx).Where(mutation.ID.Eq(id)).First()
	if err != nil {
		return nil, teamDashboardNotFound(err)
	}
	return dashboardDo, nil
}

// ListDashboards list dashboards with filter
func (r *dashboardImpl) ListDashboards(ctx context.Context, req *bo.ListDashboardReq) (*bo.ListDashboardReply, error) {
	tx, err := r.getDashboardTX(ctx)
	if err != nil {
		return nil, err
	}
	mutation := tx.Dashboard
	query := mutation.WithContext(ctx)

	if !req.Status.IsUnknown() {
		query = query.Where(mutation.Status.Eq(req.Status.GetValue()))
	}

	if req.PaginationRequest != nil {
		total, err := query.Count()
		if err != nil {
			return nil, err
		}
		req.WithTotal(total)
		query = query.Offset(req.Offset()).Limit(int(req.Limit))
	}

	dashboards, err := query.Find()
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
	tx, err := r.getDashboardTX(ctx)
	if err != nil {
		return err
	}
	mutation := tx.Dashboard
	_, err = mutation.WithContext(ctx).Where(mutation.ID.In(ids...)).UpdateColumnSimple(mutation.Status.Value(int8(status)))
	return err
}
