package biz

import (
	"context"

	"github.com/google/wire"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/do"
)

// ProviderSetBiz is biz providers.
var ProviderSetBiz = wire.NewSet(
	NewAuthBiz,
	NewPermissionBiz,
	NewResourceBiz,
	NewUserBiz,
	NewDashboardBiz,
	NewServerBiz,
	NewDict,
	NewTeam,
	NewTeamHook,
	NewMessage,
	NewSystem,
	NewTeamNotice,
	NewTeamDatasource,
	NewTeamStrategy,
)

type GetUserFun func(ctx context.Context, id uint32) (do.User, error)

type GetUsersFun func(ctx context.Context, ids ...uint32) (map[uint32]do.User, error)

var GetUser GetUserFun = func(ctx context.Context, id uint32) (do.User, error) {
	panic("not implement")
}

var GetUsers GetUsersFun = func(ctx context.Context, ids ...uint32) (map[uint32]do.User, error) {
	panic("not implement")
}
