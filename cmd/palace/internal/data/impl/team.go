package impl

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	ggorm "gorm.io/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do/system"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/repository"
	"github.com/moon-monitor/moon/cmd/palace/internal/conf"
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

func (r *teamRepoImpl) FindByID(ctx context.Context, id uint32) (*system.Team, error) {
	systemQuery := r.Team
	teamDo, err := systemQuery.WithContext(ctx).Where(systemQuery.ID.Eq(id)).First()
	if err != nil {
		return nil, teamNotFound(err)
	}
	return teamDo, nil
}

func (r *teamRepoImpl) createDatabase(c *conf.Data_Database, sqlDB *sql.DB, teamID uint32) (gorm.DB, error) {
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
	if teamDo.Capacity.AllowGroup() && c.IsGroup() {
		dbName = fmt.Sprintf("%s_%d", dbName, teamID)
	}

	switch c.GetDriver() {
	case config.Database_MYSQL:
		expr := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;", dbName)
		if _, err := sqlDB.Exec(expr); err != nil {
			return nil, err
		}
	}
	dsn := c.GenDsn(dbName)

	databaseConf := &config.Database{
		Driver:       c.GetDriver(),
		Dsn:          dsn,
		Debug:        c.GetDebug(),
		UseSystemLog: c.GetUseSystemLog(),
	}
	return gorm.NewDB(databaseConf)
}
