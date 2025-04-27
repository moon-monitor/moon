package service

import (
	"context"

	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

type TeamLogService struct {
	palace.UnimplementedTeamLogServer
}

func NewTeamLogService() *TeamLogService {
	return &TeamLogService{}
}

func (s *TeamLogService) SendMessageLogs(ctx context.Context, req *palace.SendMessageLogsRequest) (*palace.SendMessageLogsReply, error) {
	params := build.ToListSendMessageLogParams(req)
	params, err := params.WithTeamID(ctx)
	if err != nil {
		return nil, err
	}
	return &palace.SendMessageLogsReply{}, nil
}

func (s *TeamLogService) GetSendMessageLog(ctx context.Context, req *palace.OperateOneSendMessageRequest) (*common.SendMessageLog, error) {
	return &common.SendMessageLog{}, nil
}

func (s *TeamLogService) RetrySendMessage(ctx context.Context, req *palace.OperateOneSendMessageRequest) (*common.EmptyReply, error) {
	return &common.EmptyReply{}, nil
}
