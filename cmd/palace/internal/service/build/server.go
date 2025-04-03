package build

import (
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/bo"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	common "github.com/moon-monitor/moon/pkg/api/common"
)

func ToServerRegisterReq(req *common.ServerRegisterRequest) *bo.ServerRegisterReq {
	if req == nil {
		return nil
	}
	return &bo.ServerRegisterReq{
		ServerType: vobj.ServerType(req.GetServerType()),
		Server:     req.GetServer(),
		Discovery:  req.GetDiscovery(),
		TeamIds:    req.GetTeamIds(),
		IsOnline:   req.GetIsOnline(),
		Uuid:       req.GetUuid(),
	}
}
