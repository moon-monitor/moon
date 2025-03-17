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
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/password"
	"github.com/moon-monitor/moon/pkg/util/template"
	"github.com/moon-monitor/moon/pkg/util/validate"
)

func NewUserRepo(bc *conf.Bootstrap, data *data.Data, logger log.Logger) repository.User {
	return &userRepoImpl{
		bc:     bc,
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.user")),
	}
}

type userRepoImpl struct {
	bc *conf.Bootstrap
	*data.Data

	helper *log.Helper
}

func userNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorUserNotFound("user not found").WithCause(err)
	}
	return err
}

func (u *userRepoImpl) CreateUserWithOAuthUser(ctx context.Context, user bo.IOAuthUser) (userDo *system.User, err error) {
	userDo, err = u.FindByEmail(ctx, user.GetEmail())
	if err == nil {
		return userDo, nil
	}
	if !merr.IsUserNotFound(err) {
		return nil, err
	}
	pass := password.New(password.GenerateRandomPassword(8))
	enValue, err := pass.EnValue()
	if err != nil {
		return nil, err
	}
	userDo = &system.User{
		BaseModel: do.BaseModel{},
		Username:  user.GetUsername(),
		Nickname:  user.GetNickname(),
		Password:  enValue,
		Email:     user.GetEmail(),
		Phone:     "",
		Remark:    user.GetRemark(),
		Avatar:    user.GetAvatar(),
		Salt:      pass.Salt(),
		Gender:    vobj.GenderUnknown,
		Position:  vobj.RoleUser,
		Status:    vobj.UserStatusNormal,
		Roles:     nil,
	}
	userMutation := systemQuery.Use(u.GetMainDB().GetDB()).User
	if err = userMutation.WithContext(ctx).Create(userDo); err != nil {
		return nil, err
	}

	if err = u.sendUserPassword(userDo, pass.PValue()); err != nil {
		return nil, err
	}
	return userDo, nil
}

func (u *userRepoImpl) FindByID(ctx context.Context, userID uint32) (*system.User, error) {
	userQuery := systemQuery.Use(u.GetMainDB().GetDB()).User
	user, err := userQuery.WithContext(ctx).Where(userQuery.ID.Eq(userID)).First()
	if err != nil {
		return nil, userNotFound(err)
	}
	return user, nil
}

func (u *userRepoImpl) FindByEmail(ctx context.Context, email string) (*system.User, error) {
	userQuery := systemQuery.Use(u.GetMainDB().GetDB()).User
	user, err := userQuery.WithContext(ctx).Where(userQuery.Email.Eq(email)).First()
	if err != nil {
		return nil, userNotFound(err)
	}
	return user, nil
}

func (u *userRepoImpl) SetEmail(ctx context.Context, user *system.User) (*system.User, error) {
	userMutation := systemQuery.Use(u.GetMainDB().GetDB()).User
	wrapper := []gen.Condition{
		userMutation.ID.Eq(user.ID),
		userMutation.Email.Eq(""),
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
	if err = u.sendUserPassword(userDo, pass.PValue()); err != nil {
		return nil, err
	}
	return userDo, nil
}

//go:embed template/welcome.html
var welcomeEmailBody string

func (u *userRepoImpl) sendUserPassword(user *system.User, pass string) error {
	if err := validate.CheckEmail(user.Email); err != nil {
		return nil
	}

	bodyParams := map[string]string{
		"Username":    user.Email,
		"Password":    pass,
		"RedirectURI": u.bc.GetAuth().GetOauth2().GetRedirectUri(),
	}
	emailBody, err := template.HtmlFormatter(welcomeEmailBody, bodyParams)
	if err != nil {
		return err
	}
	// 发送用户密码到用户邮箱
	return u.GetEmail().SetSubject("Welcome to the Moon Monitoring System.").SetTo(user.Email).SetBody(emailBody, "text/html").Send()
}
