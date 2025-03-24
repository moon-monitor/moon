package impl

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	query "github.com/moon-monitor/moon/cmd/palace/internal/data/query/system"
)

func NewTeamRepo(d *data.Data, logger log.Logger) repository.Team {
	return &teamRepoImpl{
		Data:   d,
		helper: log.NewHelper(log.With(logger, "module", "data.repo.team")),
	}
}

type teamRepoImpl struct {
	*data.Data

	helper *log.Helper
}

func (r *teamRepoImpl) FindByID(ctx context.Context, id uint32) (*system.Team, error) {
	systemQuery := query.Use(r.GetMainDB().GetDB()).Team
	team, err := systemQuery.WithContext(ctx).Where(systemQuery.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return team, nil
}
