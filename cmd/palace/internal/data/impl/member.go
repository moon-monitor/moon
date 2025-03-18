package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	systemQuery "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
)

func NewMemberRepo(data *data.Data, logger log.Logger) repository.Member {
	return &memberImpl{
		Data:   data,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.member")),
	}
}

type memberImpl struct {
	*data.Data

	helper *log.Helper
}

func (m *memberImpl) FindByUserID(ctx context.Context, userID uint32) (*system.TeamMember, error) {
	memberQuery := systemQuery.Use(m.GetMainDB().GetDB()).TeamMember
	member, err := memberQuery.WithContext(ctx).Where(memberQuery.UserID.Eq(userID)).First()
	if err != nil {
		return nil, userNotFound(err)
	}
	return member, nil
}
