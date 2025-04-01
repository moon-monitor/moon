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

// SaveResource 保存资源
func (s *ResourceService) SaveResource(ctx context.Context, req *palacev1.SaveResourceRequest) (*common.EmptyReply, error) {
	// 转换为业务对象
	saveReq := &bo.SaveResourceReq{
		ID:     req.GetId(),
		Name:   req.GetName(),
		Path:   req.GetPath(),
		Module: vobj.ResourceModule(req.GetModule()),
		Domain: vobj.ResourceDomain(req.GetDomain()),
		Remark: req.GetRemark(),
		Allow:  vobj.ResourceAllow(req.GetAllow()),
	}

	// 保存资源
	err := s.resourceBiz.SaveResource(ctx, saveReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// BatchUpdateResourceStatus 批量更新资源状态
func (s *ResourceService) BatchUpdateResourceStatus(ctx context.Context, req *palacev1.BatchUpdateResourceStatusRequest) (*common.EmptyReply, error) {
	// 转换为业务对象
	updateReq := &bo.BatchUpdateResourceStatusReq{
		IDs:    req.GetIds(),
		Status: vobj.ResourceStatus(req.GetStatus()),
	}

	// 批量更新状态
	err := s.resourceBiz.BatchUpdateResourceStatus(ctx, updateReq)
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// DeleteResource 删除资源
func (s *ResourceService) DeleteResource(ctx context.Context, req *palacev1.DeleteResourceRequest) (*common.EmptyReply, error) {
	// 删除资源
	err := s.resourceBiz.DeleteResource(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &common.EmptyReply{Message: "success"}, nil
}

// GetResource 获取资源详情
func (s *ResourceService) GetResource(ctx context.Context, req *palacev1.GetResourceRequest) (*palacev1.GetResourceReply, error) {
	// 获取资源详情
	resource, err := s.resourceBiz.GetResource(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	// 转换为proto对象
	return &palacev1.GetResourceReply{
		Detail: build.ToResourceItemProto(resource),
	}, nil
}

// ListResource 查询资源列表
func (s *ResourceService) ListResource(ctx context.Context, req *palacev1.ListResourceRequest) (*palacev1.ListResourceReply, error) {
	// 转换为业务对象
	listReq := &bo.ListResourceReq{
		Keyword: req.GetKeyword(),
	}

	// 转换状态过滤条件
	if len(req.GetStatus()) > 0 {
		listReq.Statuses = make([]vobj.ResourceStatus, 0, len(req.GetStatus()))
		for _, status := range req.GetStatus() {
			listReq.Statuses = append(listReq.Statuses, vobj.ResourceStatus(status))
		}
	}

	// 转换模块过滤条件
	if len(req.GetModule()) > 0 {
		listReq.Modules = make([]vobj.ResourceModule, 0, len(req.GetModule()))
		for _, module := range req.GetModule() {
			listReq.Modules = append(listReq.Modules, vobj.ResourceModule(module))
		}
	}

	// 转换领域过滤条件
	if len(req.GetDomain()) > 0 {
		listReq.Domains = make([]vobj.ResourceDomain, 0, len(req.GetDomain()))
		for _, domain := range req.GetDomain() {
			listReq.Domains = append(listReq.Domains, vobj.ResourceDomain(domain))
		}
	}

	// 查询资源列表
	resources, err := s.resourceBiz.ListResource(ctx, listReq)
	if err != nil {
		return nil, err
	}

	// 转换为proto对象
	return &palacev1.ListResourceReply{
		Items: build.ToResourceItemProtoList(resources),
	}, nil
}
