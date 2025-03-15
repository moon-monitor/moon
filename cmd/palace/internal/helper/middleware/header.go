package middleware

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"
	"github.com/moon-monitor/moon/pkg/merr"
)

const (
	XHeaderTeamID       = "X-Team-ID"
	XHeaderTeamMemberID = "X-Team-Member-ID"
	XHeaderSysPosition  = "X-Sys-Position"
	XHeaderSysRoleID    = "X-Sys-Role"
	XHeaderTeamPosition = "X-Team-Position"
	XHeaderTeamRoleID   = "X-Team-Role"
)

func BindHeaders() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			ctx, err := withAllHeaders(ctx)
			if err != nil {
				return nil, err
			}
			return handler(ctx, req)
		}
	}
}

func withAllHeaders(ctx context.Context) (context.Context, error) {
	tr, ok := transport.FromServerContext(ctx)
	if !ok {
		return nil, merr.ErrorBadRequest("not allow request")
	}
	if teamIDStr := tr.RequestHeader().Get(XHeaderTeamID); teamIDStr != "" {
		teamID, err := strconv.ParseUint(teamIDStr, 10, 32)
		if err != nil {
			return nil, merr.ErrorBadRequest("not allow request, header [%s] err", XHeaderTeamID)
		}
		ctx = WithTeamIDContext(ctx, uint32(teamID))
	}
	if teamMemberIDStr := tr.RequestHeader().Get(XHeaderTeamMemberID); teamMemberIDStr != "" {
		teamMemberID, err := strconv.ParseUint(teamMemberIDStr, 10, 32)
		if err != nil {
			return nil, merr.ErrorBadRequest("not allow request, header [%s] err", XHeaderTeamMemberID)
		}
		ctx = WithTeamMemberIDContext(ctx, uint32(teamMemberID))
	}
	if sysPositionStr := tr.RequestHeader().Get(XHeaderSysPosition); sysPositionStr != "" {
		sysPosition, err := strconv.ParseInt(sysPositionStr, 10, 32)
		if err != nil {
			return nil, merr.ErrorBadRequest("not allow request, header [%s] err", XHeaderSysPosition)
		}
		ctx = WithSysPositionContext(ctx, vobj.Role(sysPosition))
	}
	if sysRoleStr := tr.RequestHeader().Get(XHeaderSysRoleID); sysRoleStr != "" {
		sysRole, err := strconv.ParseUint(sysRoleStr, 10, 32)
		if err != nil {
			return nil, merr.ErrorBadRequest("not allow request, header [%s] err", XHeaderSysRoleID)
		}
		ctx = WithSysRoleIDContext(ctx, uint32(sysRole))
	}
	if teamPositionStr := tr.RequestHeader().Get(XHeaderTeamPosition); teamPositionStr != "" {
		teamPosition, err := strconv.ParseInt(teamPositionStr, 10, 32)
		if err != nil {
			return nil, merr.ErrorBadRequest("not allow request, header [%s] err", XHeaderTeamPosition)
		}
		ctx = WithTeamPositionContext(ctx, vobj.Role(teamPosition))
	}
	if teamRoleStr := tr.RequestHeader().Get(XHeaderTeamRoleID); teamRoleStr != "" {
		teamRole, err := strconv.ParseUint(teamRoleStr, 10, 32)
		if err != nil {
			return nil, merr.ErrorBadRequest("not allow request, header [%s] err", XHeaderTeamRoleID)
		}
		ctx = WithTeamRoleIDContext(ctx, uint32(teamRole))
	}
	return ctx, nil
}
