package impl

import (
	"context"

	"github.com/google/wire"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"

	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/teamgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
)

// ProviderSetImpl is a set of providers.
var ProviderSetImpl = wire.NewSet(
	NewUserRepo,
	NewMemberRepo,
	NewCaptchaRepo,
	NewCacheRepo,
	NewOAuthRepo,
	NewResourceRepo,
	NewTransaction,
	NewTeamRepo,
	NewDashboardRepo,
	NewDashboardChartRepo,
	NewServerRepo,
	NewTeamDictRepo,
	NewTeamHook,
)

type BizDB interface {
	GetBizDB(teamID uint32) (gorm.DB, error)
}

func getTeamBizQuery(ctx context.Context, b BizDB) (*teamgen.Query, uint32, error) {
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return nil, 0, merr.ErrorPermissionDenied("team id not found")
	}
	bizDB, err := b.GetBizDB(teamID)
	if err != nil {
		return nil, 0, err
	}
	return teamgen.Use(bizDB.GetDB()), teamID, nil
}
