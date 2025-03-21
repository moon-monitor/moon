package hook

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/httpx"
)

var _ Hook = (*wechatHook)(nil)

func NewWechatHook(api string, opts ...WechatHookOption) *wechatHook {
	h := &wechatHook{api: api}
	for _, opt := range opts {
		opt(h)
	}
	if h.helper == nil {
		WithWechatLogger(log.DefaultLogger)(h)
	}
	return h
}

func WithWechatLogger(logger log.Logger) WechatHookOption {
	return func(h *wechatHook) {
		h.helper = log.NewHelper(log.With(logger, "module", "plugin.hook.wechat"))
	}
}

type WechatHookOption func(*wechatHook)

type wechatHook struct {
	api    string
	helper *log.Helper
}

type wechatHookResp struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func (l *wechatHookResp) Error() error {
	if l.ErrCode == 0 {
		return nil
	}
	return merr.ErrorBadRequest("errcode: %d, errmsg: %s", l.ErrCode, l.ErrMsg)
}

func (h *wechatHook) Send(ctx context.Context, message Message) error {
	response, err := httpx.PostJson(ctx, h.api, message)
	if err != nil {
		h.helper.Debugf("send wechat hook failed: %v", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		h.helper.Debugf("send wechat hook failed: status code: %d", response.StatusCode)
		return merr.ErrorBadRequest("status code: %d", response.StatusCode)
	}

	var resp wechatHookResp
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		h.helper.Debugf("send wechat hook failed: %v", err)
		return err
	}

	return resp.Error()
}
