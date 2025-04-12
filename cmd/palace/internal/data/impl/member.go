package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/teamgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"

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
	teamId, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return nil, merr.ErrorPermissionDenied("team id is invalid")
	}
	bizDB, err := m.GetBizDB(teamId)
	if err != nil {
		return nil, err
	}
	memberQuery := teamgen.Use(bizDB.GetDB()).Member
	member, err := memberQuery.WithContext(ctx).Where(memberQuery.UserID.Eq(userID)).First()
	if err != nil {
		return nil, teamMemberNotFound(err)
	}
	return member, nil
}
