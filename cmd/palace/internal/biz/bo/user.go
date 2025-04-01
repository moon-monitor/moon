package bo

import "github.com/moon-monitor/moon/cmd/palace/internal/biz/vobj"

type UserUpdateInfo struct {
	Nickname string
	Avatar   string
	Gender   vobj.Gender
}
