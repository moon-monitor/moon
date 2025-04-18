package impl

import (
	"context"
	"fmt"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gen"
	"gorm.io/gen/field"
	ggorm "gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/cmd/palace/internal/data"
	"github.com/moon-monitor/moon/pkg/config"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
	"github.com/moon-monitor/moon/pkg/util/slices"
	"github.com/moon-monitor/moon/pkg/util/validate"
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

func (r *teamRepoImpl) Create(ctx context.Context, team do.Team) (do.Team, error) {
	teamMutation := getMainQuery(ctx, r).Team
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
	}
	teamDo.WithContext(ctx)
	if err := teamMutation.WithContext(ctx).Create(teamDo); err != nil {
		return nil, err
	}
	return teamDo, nil
}

func (r *teamRepoImpl) Update(ctx context.Context, team do.Team) (do.Team, error) {
	teamMutation := getMainQuery(ctx, r).Team
	wrappers := []gen.Condition{
		teamMutation.ID.Eq(team.GetID()),
	}
	mutations := []field.AssignExpr{
		teamMutation.Name.Value(team.GetName()),
		teamMutation.Remark.Value(team.GetRemark()),
		teamMutation.Logo.Value(team.GetLogo()),
	}
	_, err := teamMutation.WithContext(ctx).Where(wrappers...).UpdateColumnSimple(mutations...)
	if err != nil {
		return nil, err
	}
	return r.FindByID(ctx, team.GetID())
}

func (r *teamRepoImpl) Delete(ctx context.Context, id uint32) error {
	teamMutation := getMainQuery(ctx, r).Team
	wrappers := []gen.Condition{
		teamMutation.ID.Eq(id),
	}
	_, err := teamMutation.WithContext(ctx).Where(wrappers...).Delete()
	return err
}

func (r *teamRepoImpl) FindByID(ctx context.Context, id uint32) (do.Team, error) {
	systemQuery := getMainQuery(ctx, r).Team
	teamDo, err := systemQuery.WithContext(ctx).Where(systemQuery.ID.Eq(id)).First()
	if err != nil {
		return nil, teamNotFound(err)
	}
	return teamDo, nil
}

func (r *teamRepoImpl) List(ctx context.Context, req *bo.TeamListRequest) (*bo.TeamListReply, error) {
	query := getMainQuery(ctx, r)
	teamQuery := query.Team
	wrapper := teamQuery.WithContext(ctx)
	if !validate.TextIsNull(req.Keyword) {
		wrapper = wrapper.Where(teamQuery.Name.Like(req.Keyword))
	}
	if len(req.Status) > 0 {
		status := slices.Map(req.Status, func(statusItem vobj.TeamStatus) int8 { return statusItem.GetValue() })
		wrapper = wrapper.Where(teamQuery.Status.In(status...))
	}
	if len(req.UserIds) > 0 {
		userQuery := query.User
		users, err := userQuery.WithContext(ctx).Where(userQuery.ID.In(req.UserIds...)).Preload(userQuery.Teams).Find()
		if err != nil {
			return nil, err
		}
		if len(users) > 0 {
			var teamIds []uint32
			for _, user := range users {
				teamIds = append(teamIds, slices.Map(user.GetTeams(), func(team do.Team) uint32 { return team.GetID() })...)
			}
			if len(teamIds) > 0 {
				wrapper = wrapper.Where(teamQuery.ID.In(teamIds...))
			}
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
	return req.ToTeamListReply(teamDos), nil
}

func (r *teamRepoImpl) createDatabase(ctx context.Context, c *config.Database, teamID uint32) (gorm.DB, error) {
	teamQuery := getMainQuery(ctx, r).Team
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
