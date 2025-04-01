package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz"
	"github.com/moon-monitor/moon/cmd/palace/internal/service/build"
	"github.com/moon-monitor/moon/pkg/api/palace"
	"github.com/moon-monitor/moon/pkg/api/palace/common"
)

// UserService is a user service implementation.
type UserService struct {
	palace.UnimplementedUserServer

	userBiz *biz.UserBiz
	log     *log.Helper
}

// NewUserService creates a new user service.
func NewUserService(userBiz *biz.UserBiz, logger log.Logger) *UserService {
	return &UserService{
		userBiz: userBiz,
		log:     log.NewHelper(log.With(logger, "module", "service/user")),
	}
}

// SelfInfo retrieves the current user's information.
func (s *UserService) SelfInfo(ctx context.Context, req *common.EmptyRequest) (*palace.SelfInfoReply, error) {
	user, err := s.userBiz.GetSelfInfo(ctx)
	if err != nil {
		return nil, err
	}

	return &palace.SelfInfoReply{
		User: build.UserToUserItemProto(user),
	}, nil
}

// UpdateSelfInfo updates the current user's information.
func (s *UserService) UpdateSelfInfo(ctx context.Context, req *palace.UpdateSelfInfoRequest) (*common.EmptyReply, error) {
	// TODO: implement the logic
	return &common.EmptyReply{}, nil
}

// UpdateSelfPassword updates the current user's password.
func (s *UserService) UpdateSelfPassword(ctx context.Context, req *palace.UpdateSelfPasswordRequest) (*common.EmptyReply, error) {
	// TODO: implement the logic
	return &common.EmptyReply{}, nil
}

// LeaveTeam allows the current user to leave a team.
func (s *UserService) LeaveTeam(ctx context.Context, req *palace.LeaveTeamRequest) (*common.EmptyReply, error) {
	// TODO: implement the logic
	return &common.EmptyReply{}, nil
}

// JoinTeam allows the current user to join a team.
func (s *UserService) JoinTeam(ctx context.Context, req *palace.JoinTeamRequest) (*common.EmptyReply, error) {
	// TODO: implement the logic
	return &common.EmptyReply{}, nil
}

// CreateTeam allows the current user to create a new team.
func (s *UserService) CreateTeam(ctx context.Context, req *palace.CreateTeamRequest) (*common.EmptyReply, error) {
	// TODO: implement the logic
	return &common.EmptyReply{}, nil
}

// SelfTeamList retrieves the list of teams the current user is a member of.
func (s *UserService) SelfTeamList(ctx context.Context, req *common.EmptyRequest) (*palace.SelfTeamListReply, error) {
	// TODO: implement the logic
	return &palace.SelfTeamListReply{}, nil
}

// SelfSubscribeTeamStrategies retrieves the list of team strategies the current user is subscribed to.
func (s *UserService) SelfSubscribeTeamStrategies(ctx context.Context, req *palace.SelfSubscribeTeamStrategiesRequest) (*palace.SelfSubscribeTeamStrategiesReply, error) {
	// TODO: implement the logic
	return &palace.SelfSubscribeTeamStrategiesReply{}, nil
}
