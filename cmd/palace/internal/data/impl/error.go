package impl

import (
	"github.com/go-kratos/kratos/v2/errors"
	"gorm.io/gorm"

	"github.com/moon-monitor/moon/pkg/merr"
)

func userNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorUserNotFound("user not found").WithCause(err)
	}
	return err
}

func oauthUserNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorUserNotFound("oauth user not found").WithCause(err)
	}
	return err
}

func teamDashboardNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorNotFound("team dashboard not found").WithCause(err)
	}
	return err
}

func teamDashboardChartNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorNotFound("team dashboard chart not found").WithCause(err)
	}
	return err
}

func teamNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorNotFound("team not found").WithCause(err)
	}
	return err
}

func teamMemberNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorNotFound("team member not found").WithCause(err)
	}
	return err
}

func resourceNotFound(err error) error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return merr.ErrorNotFound("resource not found").WithCause(err)
	}
	return err
}
