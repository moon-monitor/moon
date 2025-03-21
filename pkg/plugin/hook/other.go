package hook

import (
	"context"
	"io"
	"net/http"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/httpx"
)

var _ Sender = (*otherHook)(nil)

func NewOtherHook(api string, opts ...OtherHookOption) Sender {
	h := &otherHook{
		api: api,
	}
	for _, opt := range opts {
		opt(h)
	}
	if h.helper == nil {
		WithOtherLogger(log.DefaultLogger)(h)
	}
	return h
}

func WithOtherBasicAuth(username, password string) OtherHookOption {
	return func(h *otherHook) {
		h.basicAuth = &BasicAuth{
			Username: username,
			Password: password,
		}
	}
}

func WithOtherHeader(header http.Header) OtherHookOption {
	return func(h *otherHook) {
		h.header = header
	}
}

func WithOtherLogger(logger log.Logger) OtherHookOption {
	return func(h *otherHook) {
		h.helper = log.NewHelper(log.With(logger, "module", "plugin.hook.other"))
	}
}

type OtherHookOption func(*otherHook)

type otherHook struct {
	api       string
	basicAuth *BasicAuth
	header    http.Header
	helper    *log.Helper
}

// Send implements Hook.
func (o *otherHook) Send(ctx context.Context, message Message) (err error) {
	defer func() {
		if err != nil {
			o.helper.Warnw("msg", "send other hook failed", "error", err, "req", string(message))
		}
	}()
	response, err := httpx.PostJson(ctx, o.api, []byte(message))
	if err != nil {
		o.helper.Warnf("send other hook failed: %v", err)
		return err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		o.helper.Warnf("read other hook response body failed: %v", err)
		return merr.ErrorBadRequest("read other hook response body failed: %v", err)
	}
	if response.StatusCode != http.StatusOK {
		o.helper.Warnf("send other hook failed: status code: %d, response: %s", response.StatusCode, string(body))
		return merr.ErrorBadRequest("status code: %d, response: %s", response.StatusCode, string(body))
	}

	return nil
}
