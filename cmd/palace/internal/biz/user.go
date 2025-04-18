package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/password"
)

// UserBiz is a user business logic implementation.
type UserBiz struct {
	userRepo repository.User
	log      *log.Helper
}

// NewUserBiz creates a new UserBiz instance.
func NewUserBiz(userRepo repository.User, logger log.Logger) *UserBiz {
	return &UserBiz{
		userRepo: userRepo,
		log:      log.NewHelper(log.With(logger, "module", "biz.user")),
	}
}

// GetSelfInfo retrieves the current user's information from the context.
func (b *UserBiz) GetSelfInfo(ctx context.Context) (do.User, error) {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("user not found in context")
	}

	user, err := b.userRepo.FindByID(ctx, userID)
	if err != nil {
		return nil, merr.ErrorInternalServerError("failed to find user").WithCause(err)
	}

	if user == nil {
		return nil, merr.ErrorUserNotFound("user not found")
	}

	return user, nil
}

// UpdateSelfInfo updates the current user's profile information.
func (b *UserBiz) UpdateSelfInfo(ctx context.Context, userUpdateInfo *bo.UserUpdateInfo) error {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorUnauthorized("user not found in context")
	}

	user, err := b.userRepo.FindByID(ctx, userID)
	if err != nil {
		return merr.ErrorInternalServerError("failed to find user").WithCause(err)
	}

	if user == nil {
		return merr.ErrorUserNotFound("user not found")
	}

	if err = b.userRepo.UpdateUserInfo(ctx, userUpdateInfo.WithUser(user)); err != nil {
		return merr.ErrorInternalServerError("failed to update user info").WithCause(err)
	}

	return nil
}

func (b *UserBiz) UpdateUserBaseInfo(ctx context.Context, userUpdateInfo *bo.UserUpdateInfo) error {
	user, err := b.userRepo.FindByID(ctx, userUpdateInfo.GetUserID())
	if err != nil {
		return err
	}
	if err = b.userRepo.UpdateUserInfo(ctx, userUpdateInfo.WithUser(user)); err != nil {
		return merr.ErrorInternalServerError("failed to update user info").WithCause(err)
	}
	return nil
}

// UpdateSelfPassword updates the current user's password
func (b *UserBiz) UpdateSelfPassword(ctx context.Context, passwordUpdateInfo *bo.PasswordUpdateInfo) error {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return merr.ErrorUnauthorized("user not found in context")
	}

	user, err := b.userRepo.FindByID(ctx, userID)
	if err != nil {
		return err
	}

	if user == nil {
		return merr.ErrorUserNotFound("user not found")
	}

	// Verify old password
	if !user.ValidatePassword(passwordUpdateInfo.OldPassword) {
		return merr.ErrorPasswordError("old password is incorrect")
	}

	// Generate new password
	pass := password.New(passwordUpdateInfo.NewPassword)
	encryptedPassword, err := pass.EnValue()
	if err != nil {
		return merr.ErrorInternalServerError("failed to encrypt password").WithCause(err)
	}

	updateUserPasswordInfo := &bo.UpdateUserPasswordInfo{
		UserID:         userID,
		Password:       encryptedPassword,
		Salt:           pass.Salt(),
		OriginPassword: passwordUpdateInfo.NewPassword,
		SendEmailFun:   passwordUpdateInfo.SendEmailFun,
	}
	// Update password through repository
	if err = b.userRepo.UpdatePassword(ctx, updateUserPasswordInfo); err != nil {
		return merr.ErrorInternalServerError("failed to update password").WithCause(err)
	}

	return nil
}

// GetUserTeams retrieves all teams that the user is a member of
func (b *UserBiz) GetUserTeams(ctx context.Context) ([]do.Team, error) {
	userID, ok := permission.GetUserIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorUnauthorized("user not found in context")
	}

	teams, err := b.userRepo.GetTeamsByUserID(ctx, userID)
	if err != nil {
		return nil, merr.ErrorInternalServerError("failed to get user teams").WithCause(err)
	}

	return teams, nil
}

func (b *UserBiz) UpdateUserStatus(ctx context.Context, req *bo.UpdateUserStatusRequest) error {
	return b.userRepo.UpdateUserStatus(ctx, req)
}

func (b *UserBiz) ResetUserPassword(ctx context.Context, req *bo.ResetUserPasswordRequest) error {
	user, err := b.userRepo.FindByID(ctx, req.UserId)
	if err != nil {
		return err
	}
	newPass := password.GenerateRandomPassword(8)
	pass := password.New(newPass)
	enValue, err := pass.EnValue()
	if err != nil {
		return err
	}
	updateUserPasswordInfo := &bo.UpdateUserPasswordInfo{
		UserID:         user.GetID(),
		Password:       enValue,
		Salt:           pass.Salt(),
		OriginPassword: newPass,
		SendEmailFun:   req.SendEmailFun,
	}
	return b.userRepo.UpdatePassword(ctx, updateUserPasswordInfo)
}
