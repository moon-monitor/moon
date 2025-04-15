package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	palacev1 "github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type DashboardService struct {
	palacev1.UnimplementedTeamDashboardServer

	dashboardBiz *biz.DashboardBiz
	helper       *log.Helper
}

func NewDashboardService(dashboardBiz *biz.DashboardBiz, logger log.Logger) *DashboardService {
	return &DashboardService{
		dashboardBiz: dashboardBiz,
		helper:       log.NewHelper(log.With(logger, "module", "service.dashboard")),
	}
}

// SaveDashboard save dashboard
func (s *DashboardService) SaveDashboard(ctx context.Context, req *palacev1.SaveTeamDashboardRequest) (*common.EmptyReply, error) {
	// Convert to business object
	saveReq := &bo.SaveDashboardReq{
		ID:       req.GetDashboardID(),
		Title:    req.GetTitle(),
		Remark:   req.GetRemark(),
		Status:   vobj.GlobalStatus(req.GetStatus()),
		ColorHex: req.GetColorHex(),
	}

	// Save dashboard
	err := s.dashboardBiz.SaveDashboard(ctx, saveReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// DeleteDashboard delete dashboard
func (s *DashboardService) DeleteDashboard(ctx context.Context, req *palacev1.DeleteTeamDashboardRequest) (*common.EmptyReply, error) {
	// Delete dashboard
	err := s.dashboardBiz.DeleteDashboard(ctx, req.GetDashboardID())
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// GetDashboard get dashboard
func (s *DashboardService) GetDashboard(ctx context.Context, req *palacev1.GetTeamDashboardRequest) (*palacev1.GetTeamDashboardReply, error) {
	// Get dashboard
	dashboard, err := s.dashboardBiz.GetDashboard(ctx, req.GetDashboardID())
	if err != nil {
		return nil, err
	}

	// Convert to proto object
	return &palacev1.GetTeamDashboardReply{
		Dashboard: build.ToDashboardItemProto(dashboard),
	}, nil
}

// ListDashboard list dashboard
func (s *DashboardService) ListDashboard(ctx context.Context, req *palacev1.ListTeamDashboardRequest) (*palacev1.ListTeamDashboardReply, error) {
	// Convert to business object
	listReq := &bo.ListDashboardReq{
		PaginationRequest: build.ToPaginationRequest(req.GetPagination()),
	}

	// List dashboard
	dashboards, err := s.dashboardBiz.ListDashboard(ctx, listReq)
	if err != nil {
		return nil, err
	}

	// Convert to proto object
	return &palacev1.ListTeamDashboardReply{
		Items:      build.ToDashboardItemProtoList(dashboards.Items),
		Pagination: build.ToPaginationReplyProto(dashboards.PaginationReply),
	}, nil
}

// BatchUpdateDashboardStatus batch update dashboard status
func (s *DashboardService) BatchUpdateDashboardStatus(ctx context.Context, req *palacev1.UpdateTeamDashboardStatusRequest) (*common.EmptyReply, error) {
	// Convert to business object
	updateReq := &bo.BatchUpdateDashboardStatusReq{
		Ids:    req.GetDashboardIds(),
		Status: vobj.GlobalStatus(req.GetStatus()),
	}

	// Batch update status
	err := s.dashboardBiz.BatchUpdateDashboardStatus(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// SaveDashboardChart save dashboard chart
func (s *DashboardService) SaveDashboardChart(ctx context.Context, req *palacev1.SaveTeamDashboardChartRequest) (*common.EmptyReply, error) {
	// Convert to business object
	saveReq := &bo.SaveDashboardChartReq{
		ID:          req.GetChartID(),
		DashboardID: req.GetDashboardID(),
		Title:       req.GetTitle(),
		Remark:      req.GetRemark(),
		Status:      vobj.GlobalStatus(req.GetStatus()),
		Url:         req.GetUrl(),
		Width:       req.GetWidth(),
		Height:      req.GetHeight(),
	}

	// Save dashboard chart
	err := s.dashboardBiz.SaveDashboardChart(ctx, saveReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// DeleteDashboardChart delete dashboard chart
func (s *DashboardService) DeleteDashboardChart(ctx context.Context, req *palacev1.DeleteTeamDashboardChartRequest) (*common.EmptyReply, error) {
	// Delete dashboard chart
	err := s.dashboardBiz.DeleteDashboardChart(ctx, req.GetChartID())
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// GetDashboardChart get dashboard chart
func (s *DashboardService) GetDashboardChart(ctx context.Context, req *palacev1.GetTeamDashboardChartRequest) (*palacev1.GetTeamDashboardChartReply, error) {
	// Get dashboard chart
	chart, err := s.dashboardBiz.GetDashboardChart(ctx, req.GetChartID())
	if err != nil {
		return nil, err
	}

	// Convert to proto object
	return &palacev1.GetTeamDashboardChartReply{
		Chart: build.ToDashboardChartItemProto(chart),
	}, nil
}

// ListDashboardChart list dashboard chart
func (s *DashboardService) ListDashboardChart(ctx context.Context, req *palacev1.ListTeamDashboardChartRequest) (*palacev1.ListTeamDashboardChartReply, error) {
	// Convert to business object
	listReq := &bo.ListDashboardChartReq{
		PaginationRequest: build.ToPaginationRequest(req.GetPagination()),
		DashboardID:       req.GetDashboardID(),
	}

	// List dashboard chart
	charts, err := s.dashboardBiz.ListDashboardChart(ctx, listReq)
	if err != nil {
		return nil, err
	}

	// Convert to proto object
	return &palacev1.ListTeamDashboardChartReply{
		Items:      build.ToDashboardChartItemProtoList(charts.Items),
		Pagination: build.ToPaginationReplyProto(charts.PaginationReply),
	}, nil
}

// BatchUpdateDashboardChartStatus batch update dashboard chart status
func (s *DashboardService) BatchUpdateDashboardChartStatus(ctx context.Context, req *palacev1.UpdateTeamDashboardChartStatusRequest) (*common.EmptyReply, error) {
	// Convert to business object
	updateReq := &bo.BatchUpdateDashboardChartStatusReq{
		Ids:    req.GetChartIds(),
		Status: vobj.GlobalStatus(req.GetStatus()),
	}

	// Batch update status
	err := s.dashboardBiz.BatchUpdateDashboardChartStatus(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}
