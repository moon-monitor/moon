package service

import (
	"context"

	"github.com/moon-monitor/moon/pkg/api/common"
	pb "github.com/moon-monitor/moon/pkg/api/houyi/v1"
)

type QueryService struct {
	pb.UnimplementedQueryServer
}

func NewQueryService() *QueryService {
	return &QueryService{}
}

func (s *QueryService) MetricDatasourceQuery(ctx context.Context, req *pb.MetricDatasourceQueryRequest) (*common.MetricDatasourceQueryReply, error) {
	return &common.MetricDatasourceQueryReply{}, nil
}
