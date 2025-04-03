package impl

import (
	"context"
	_ "embed"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/crypto"
	"github.com/moon-monitor/moon/pkg/util/password"
	"github.com/moon-monitor/moon/pkg/util/template"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewUserRepo(bc *conf.Bootstrap, data *data.Data, logger log.Logger) repository.User {
	return &userRepoImpl{
		bc:     bc,
		Data:   data,
		Query:  systemgen.Use(data.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.user")),
	}
}

type userRepoImpl struct {
	bc *conf.Bootstrap
	*data.Data
	*systemgen.Query
	helper *log.Helper
}

func (u *userRepoImpl) CreateUserWithOAuthUser(ctx context.Context, user bo.IOAuthUser, sendEmailFunc repository.SendEmailFunc) (userDo *system.User, err error) {
	userDo, err = u.FindByEmail(ctx, crypto.String(user.GetEmail()))
	if err == nil {
		return userDo, nil
	}
	if !merr.IsUserNotFound(err) {
		return nil, err
	}
	userDo = &system.User{
		BaseModel: do.BaseModel{},
		Username:  user.GetUsername(),
		Nickname:  user.GetNickname(),
		Password:  "",
		Email:     crypto.String(user.GetEmail()),
		Phone:     "",
		Remark:    user.GetRemark(),
		Avatar:    user.GetAvatar(),
		Salt:      "",
		Gender:    vobj.GenderUnknown,
		Position:  vobj.RoleUser,
		Status:    vobj.UserStatusNormal,
		Roles:     nil,
	}
	return u.Create(ctx, userDo, sendEmailFunc)
}

func (u *userRepoImpl) Create(ctx context.Context, user *system.User, sendEmailFunc repository.SendEmailFunc) (*system.User, error) {
	pass := password.New(password.GenerateRandomPassword(8))
	enValue, err := pass.EnValue()
	if err != nil {
		return nil, err
	}
	user.Password = enValue
	user.Salt = pass.Salt()
	userMutation := u.User
	if err = userMutation.WithContext(ctx).Create(user); err != nil {
		return nil, err
	}
	if err = u.sendUserPassword(ctx, user, pass.PValue(), sendEmailFunc); err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userRepoImpl) FindByID(ctx context.Context, userID uint32) (*system.User, error) {
	userQuery := u.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(userID)).Preload(userQuery.Roles.Resources).First()
	if err != nil {
		return nil, userNotFound(err)
	}
	return user, nil
}

func (u *userRepoImpl) FindByEmail(ctx context.Context, email crypto.String) (*system.User, error) {
	userQuery := u.User
	user, err := userQuery.WithContext(ctx).Where(userQuery.Email.Eq(email)).First()
	if err != nil {
		return nil, userNotFound(err)
	}
	return user, nil
}

func (u *userRepoImpl) SetEmail(ctx context.Context, user *system.User, sendEmailFunc repository.SendEmailFunc) (*system.User, error) {
	userMutation := u.User
	wrapper := []gen.Condition{
		userMutation.ID.Eq(user.ID),
		userMutation.Email.Eq(crypto.String("")),
	}
	pass := password.New(password.GenerateRandomPassword(8))
	enValue, err := pass.EnValue()
	if err != nil {
		return nil, err
	}
	mutations := []field.AssignExpr{
		userMutation.Email.Value(user.Email),
		userMutation.Password.Value(enValue),
		userMutation.Salt.Value(pass.Salt()),
	}
	result, err := userMutation.WithContext(ctx).Where(wrapper...).UpdateSimple(mutations...)
	if err != nil {
		return nil, err
	}
	if result.RowsAffected == 0 {
		return nil, merr.ErrorUserNotFound("user not found")
	}
	userDo, err := userMutation.WithContext(ctx).Where(userMutation.ID.Eq(user.ID), userMutation.Email.Eq(user.Email)).First()
	if err != nil {
		return nil, userNotFound(err)
	}
	if err = u.sendUserPassword(ctx, userDo, pass.PValue(), sendEmailFunc); err != nil {
		return nil, err
	}
	return userDo, nil
}

//go:embed template/welcome.html
var welcomeEmailBody string

func (u *userRepoImpl) sendUserPassword(ctx context.Context, user *system.User, pass string, sendEmailFunc repository.SendEmailFunc) error {
	if err := validate.CheckEmail(string(user.Email)); err != nil {
		return nil
	}

	bodyParams := map[string]string{
		"Username":    string(user.Email),
		"Password":    pass,
		"RedirectURI": u.bc.GetAuth().GetOauth2().GetRedirectUri(),
	}
	emailBody, err := template.HtmlFormatter(welcomeEmailBody, bodyParams)
	if err != nil {
		return err
	}
	sendEmailParams := &bo.SendEmailParams{
		Email:       string(user.Email),
		Body:        emailBody,
		Subject:     "Welcome to the Moon Monitoring System.",
		ContentType: "text/html",
	}
	// send email to user
	return sendEmailFunc(ctx, sendEmailParams)
}

// GetTeamsByUserID 获取用户所属的所有团队
func (u *userRepoImpl) GetTeamsByUserID(ctx context.Context, userID uint32) ([]*system.Team, error) {
	// 获取用户加入的所有团队
	teamMemberQuery := u.TeamMember
	teamQuery := u.Team

	// 查找用户的所有团队成员记录
	members, err := teamMemberQuery.WithContext(ctx).
		Where(teamMemberQuery.UserID.Eq(userID), teamMemberQuery.Status.Eq(int8(vobj.MemberStatusNormal))).
		Find()
	if err != nil {
		return nil, err
	}

	if len(members) == 0 {
		return []*system.Team{}, nil
	}

	// 提取所有团队ID
	teamIDs := make([]uint32, 0, len(members))
	for _, member := range members {
		teamIDs = append(teamIDs, member.TeamID)
	}

	// 查询所有团队信息，并预加载Leader和Creator关系
	teams, err := teamQuery.WithContext(ctx).
		Where(teamQuery.ID.In(teamIDs...), teamQuery.Status.Eq(int8(vobj.TeamStatusNormal))).
		Preload(teamQuery.Leader).
		Find()
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// GetMemberByUserIDAndTeamID 获取用户在特定团队中的成员信息
func (u *userRepoImpl) GetMemberByUserIDAndTeamID(ctx context.Context, userID, teamID uint32) (*system.TeamMember, error) {
	teamMemberQuery := u.TeamMember

	// 查找用户在指定团队中的成员记录，包括角色信息
	member, err := teamMemberQuery.WithContext(ctx).
		Where(teamMemberQuery.UserID.Eq(userID), teamMemberQuery.TeamID.Eq(teamID)).
		Preload(teamMemberQuery.Roles.Resources).
		First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, merr.ErrorPermissionDenied("team member not found")
		}
		return nil, err
	}

	return member, nil
}

// GetAllTeamMembers 获取用户所有团队的成员信息
func (u *userRepoImpl) GetAllTeamMembers(ctx context.Context, userID uint32) ([]*system.TeamMember, error) {
	teamMemberQuery := u.TeamMember

	// 查找用户的所有团队成员记录，包括角色信息
	members, err := teamMemberQuery.WithContext(ctx).
		Where(teamMemberQuery.UserID.Eq(userID)).
		Preload(teamMemberQuery.Roles.Resources).
		Find()
	if err != nil {
		return nil, err
	}

	return members, nil
}

// GetTeamsByIDs 根据团队ID列表获取团队
func (u *userRepoImpl) GetTeamsByIDs(ctx context.Context, teamIDs []uint32) ([]*system.Team, error) {
	// 查询所有团队信息
	teamQuery := u.Team
	teams, err := teamQuery.WithContext(ctx).
		Where(teamQuery.ID.In(teamIDs...), teamQuery.Status.Eq(int8(vobj.TeamStatusNormal))).
		Find()
	if err != nil {
		return nil, err
	}
	return teams, nil
}

// UpdateSelfInfo updates the user's profile information
func (u *userRepoImpl) UpdateSelfInfo(ctx context.Context, user *system.User) error {
	userMutation := u.User

	// Only update the relevant fields
	_, err := userMutation.WithContext(ctx).
		Where(userMutation.ID.Eq(user.ID)).
		UpdateSimple(
			userMutation.Nickname.Value(user.Nickname),
			userMutation.Avatar.Value(user.Avatar),
			userMutation.Gender.Value(int8(user.Gender)),
		)

	if err != nil {
		return err
	}

	return nil
}

// UpdatePassword updates the user's password in the database
func (u *userRepoImpl) UpdatePassword(ctx context.Context, updateUserPasswordInfo *bo.UpdateUserPasswordInfo) error {
	userMutation := u.User

	// Update password and salt fields
	_, err := userMutation.WithContext(ctx).
		Where(userMutation.ID.Eq(updateUserPasswordInfo.UserID)).
		UpdateSimple(
			userMutation.Password.Value(updateUserPasswordInfo.Password),
			userMutation.Salt.Value(updateUserPasswordInfo.Salt),
		)

	return err
}
