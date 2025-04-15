package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/team"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
)

func NewMemberRepo(data *data.Data, logger log.Logger) repository.Member {
	return &memberImpl{
		Data:   data,
		Query:  systemgen.Use(data.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.member")),
	}
}

type memberImpl struct {
	*data.Data
	*systemgen.Query
	helper *log.Helper
}

func (m *memberImpl) FindByUserID(ctx context.Context, userID uint32) (*team.Member, error) {
	query, teamID, err := getTeamBizQuery(ctx, m)
	if err != nil {
		return nil, err
	}
	memberQuery := query.Member
	wrappers := []gen.Condition{
		memberQuery.TeamID.Eq(teamID),
		memberQuery.UserID.Eq(userID),
	}
	member, err := memberQuery.WithContext(ctx).Where(wrappers...).First()
	if err != nil {
		return nil, teamMemberNotFound(err)
	}
	return member, nil
}
