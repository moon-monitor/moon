package impl

import (
	"context"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
	"gorm.io/gen"
	"gorm.io/gen/field"
	ggorm "gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
)

func NewTeamRepo(d *data.Data, logger log.Logger) repository.Team {
	return &teamRepoImpl{
		Data:   d,
		Query:  systemgen.Use(d.GetMainDB().GetDB()),
		helper: log.NewHelper(log.With(logger, "module", "data.repo.team")),
	}
}

type teamRepoImpl struct {
	*data.Data
	*systemgen.Query
	helper *log.Helper
}

func (r *teamRepoImpl) Create(ctx context.Context, team bo.Team) error {
	teamMutation := r.Team
	teamDo := &system.Team{
		Name:      team.GetName(),
		Status:    team.GetStatus(),
		Remark:    team.GetRemark(),
		Logo:      team.GetLogo(),
		LeaderID:  team.GetLeaderID(),
		UUID:      team.GetUUID(),
		Capacity:  team.GetCapacity(),
		Leader:    nil,
		Admins:    nil,
		Resources: nil,
		DBName:    "",
	}
	teamDo.WithContext(ctx)
	return teamMutation.WithContext(ctx).Create(teamDo)
}

func (r *teamRepoImpl) Update(ctx context.Context, team bo.Team) error {
	teamMutation := r.Team
	wrappers := []gen.Condition{
		teamMutation.ID.Eq(team.GetTeamID()),
	}
	mutations := []field.AssignExpr{
		teamMutation.Name.Value(team.GetName()),
		teamMutation.Remark.Value(team.GetRemark()),
		teamMutation.Logo.Value(team.GetLogo()),
	}
	_, err := teamMutation.WithContext(ctx).Where(wrappers...).UpdateColumnSimple(mutations...)
	return err
}

func (r *teamRepoImpl) Delete(ctx context.Context, id uint32) error {
	teamMutation := r.Team
	wrappers := []gen.Condition{
		teamMutation.ID.Eq(id),
	}
	_, err := teamMutation.WithContext(ctx).Where(wrappers...).Delete()
	return err
}

func (r *teamRepoImpl) FindByID(ctx context.Context, id uint32) (*system.Team, error) {
	systemQuery := r.Team
	teamDo, err := systemQuery.WithContext(ctx).Where(systemQuery.ID.Eq(id)).First()
	if err != nil {
		return nil, teamNotFound(err)
	}
	return teamDo, nil
}

func (r *teamRepoImpl) List(ctx context.Context, req *bo.TeamListRequest) (*bo.TeamListReply, error) {
	teamQuery := r.Team
	wrapper := teamQuery.WithContext(ctx)
	if !validate.TextIsNull(req.Keyword) {
		wrapper = wrapper.Where(teamQuery.Name.Like(req.Keyword))
	}
	if len(req.Status) > 0 {
		status := slices.Map(req.Status, func(statusItem vobj.TeamStatus) int8 { return statusItem.GetValue() })
		wrapper = wrapper.Where(teamQuery.Status.In(status...))
	}
	if len(req.UserIds) > 0 {
		var teamIDs []uint32
		userTeamQuery := r.UserTeam
		err := userTeamQuery.WithContext(ctx).Select(userTeamQuery.TeamID).Where(userTeamQuery.UserID.In(req.UserIds...)).Scan(&teamIDs)
		if err != nil {
			return nil, err
		}
		if len(teamIDs) > 0 {
			wrapper = wrapper.Where(teamQuery.ID.In(teamIDs...))
		}
		wrapper = wrapper.Where(teamQuery.LeaderID.In(req.UserIds...))
	}
	if validate.IsNotNil(req.PaginationRequest) {
		total, err := wrapper.Count()
		if err != nil {
			return nil, err
		}
		wrapper = wrapper.Offset(req.Offset()).Limit(int(req.Limit))
		req.WithTotal(total)
	}

	teamDos, err := wrapper.Find()
	if err != nil {
		return nil, err
	}
	return &bo.TeamListReply{
		PaginationReply: req.ToReply(),
		Items:           slices.Map(teamDos, func(teamDo *system.Team) bo.Team { return teamDo }),
	}, nil
}

func (r *teamRepoImpl) createDatabase(c *config.Database, teamID uint32) (gorm.DB, error) {
	teamQuery := r.Team
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	teamDo, err := teamQuery.WithContext(ctx).Where(teamQuery.ID.Eq(teamID)).First()
	if err != nil {
		if errors.Is(err, ggorm.ErrRecordNotFound) {
			return nil, merr.ErrorNotFound("team %d not found", teamID)
		}
		return nil, err
	}

	dbName := c.GetDbName()
	if teamDo.Capacity.AllowGroup() {
		dbName = fmt.Sprintf("%s_%d", dbName, teamID)
	}
	c.DbName = dbName
	return gorm.NewDB(c)
}
