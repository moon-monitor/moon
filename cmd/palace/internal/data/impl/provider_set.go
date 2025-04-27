package impl

import (
	"context"

	"github.com/google/wire"

	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/eventgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/systemgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/data/query/teamgen"
	"github.com/moon-monitor/moon/cmd/palace/internal/helper/permission"
	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/plugin/gorm"
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
	NewTeamConfigEmailRepo,
	NewTeamConfigSMSRepo,
	NewMenuRepo,
	NewTeamRole,
	NewRoleRepo,
	NewAuditRepo,
	NewOperateLogRepo,
	NewInviteRepo,
	NewTeamNotice,
	NewTeamMetricDatasourceRepo,
	NewTeamStrategyGroupRepo,
	NewTeamStrategyRepo,
	NewTeamStrategyMetricRepo,
	NewSendMessageLog,
)

type MainDB interface {
	GetMainDB() gorm.DB
}

type BizDB interface {
	GetBizDB(teamID uint32) (gorm.DB, error)
}

type EventDB interface {
	GetEventDB(teamID uint32) (gorm.DB, error)
}

func getMainQuery(ctx context.Context, m MainDB) *systemgen.Query {
	db := GetMainDBTransaction(ctx, m)
	return systemgen.Use(db)
}

func getTeamBizQuery(ctx context.Context, b BizDB) (*teamgen.Query, uint32, error) {
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return nil, 0, merr.ErrorPermissionDenied("team id not found")
	}
	bizTranceDB, err := GetBizTransactionDB(ctx, b)
	if err != nil {
		return nil, 0, err
	}
	return teamgen.Use(bizTranceDB), teamID, nil
}

func getTeamEventQuery(ctx context.Context, e EventDB) (*eventgen.Query, uint32, error) {
	teamID, ok := permission.GetTeamIDByContext(ctx)
	if !ok {
		return nil, 0, merr.ErrorPermissionDenied("team id not found")
	}
	eventTranceDB, err := GetEventDBTransaction(ctx, e)
	if err != nil {
		return nil, 0, err
	}
	return eventgen.Use(eventTranceDB), teamID, nil
}
