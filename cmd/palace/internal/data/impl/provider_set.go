package impl

import (
	"github.com/google/wire"
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
)
