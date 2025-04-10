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

type ResourceService struct {
	palacev1.UnimplementedResourceServer

	resourceBiz *biz.ResourceBiz
	helper      *log.Helper
}

func NewResourceService(resourceBiz *biz.ResourceBiz, logger log.Logger) *ResourceService {
	return &ResourceService{
		resourceBiz: resourceBiz,
		helper:      log.NewHelper(log.With(logger, "module", "service.resource")),
	}
}

func (s *ResourceService) BatchUpdateResourceStatus(ctx context.Context, req *palacev1.BatchUpdateResourceStatusRequest) (*common.EmptyReply, error) {
	updateReq := &bo.BatchUpdateResourceStatusReq{
		IDs:    req.GetIds(),
		Status: vobj.GlobalStatus(req.GetStatus()),
	}

	err := s.resourceBiz.BatchUpdateResourceStatus(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

func (s *ResourceService) GetResource(ctx context.Context, req *palacev1.GetResourceRequest) (*palacev1.GetResourceReply, error) {
	resource, err := s.resourceBiz.GetResource(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &palacev1.GetResourceReply{
		Detail: build.ToResourceItemProto(resource),
	}, nil
}

func (s *ResourceService) ListResource(ctx context.Context, req *palacev1.ListResourceRequest) (*palacev1.ListResourceReply, error) {
	listReq := &bo.ListResourceReq{
		Statuses:          build.Statuses(req.GetStatus()),
		Keyword:           req.GetKeyword(),
		PaginationRequest: build.ToPaginationRequest(req.GetPagination()),
	}

	resourcesReply, err := s.resourceBiz.ListResource(ctx, listReq)
	if err != nil {
		return nil, err
	}

	return &palacev1.ListResourceReply{
		Items:      build.ToResourceItemProtoList(resourcesReply.Resources),
		Pagination: build.ToPaginationReplyProto(resourcesReply.PaginationReply),
	}, nil
}

func (s *ResourceService) GetResourceMenuTree(ctx context.Context, _ *common.EmptyRequest) (*palacev1.GetResourceMenuTreeReply, error) {
	menus, err := s.resourceBiz.Menus(ctx, vobj.MenuTypeMenuSystem)
	if err != nil {
		return nil, err
	}
	return &palacev1.GetResourceMenuTreeReply{
		Items: build.ToMenuTreeProto(menus),
	}, nil
}

func (s *ResourceService) GetTeamResourceMenuTree(ctx context.Context, _ *common.EmptyRequest) (*palacev1.GetResourceMenuTreeReply, error) {
	menus, err := s.resourceBiz.Menus(ctx, vobj.MenuTypeMenuTeam)
	if err != nil {
		return nil, err
	}
	return &palacev1.GetResourceMenuTreeReply{
		Items: build.ToMenuTreeProto(menus),
	}, nil
}
