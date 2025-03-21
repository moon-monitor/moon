package hook

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/httpx"
)

var _ Hook = (*feishuHook)(nil)

func NewFeishuHook(api, secret string, opts ...FeishuHookOption) *feishuHook {
	h := &feishuHook{api: api, secret: secret}
	for _, opt := range opts {
		opt(h)
	}
	if h.helper == nil {
		WithFeishuLogger(log.DefaultLogger)(h)
	}
	return h
}

func WithFeishuLogger(logger log.Logger) FeishuHookOption {
	return func(h *feishuHook) {
		h.helper = log.NewHelper(log.With(logger, "module", "plugin.hook.feishu"))
	}
}

type FeishuHookOption func(*feishuHook)

type feishuHook struct {
	api    string
	secret string
	helper *log.Helper
}

type feishuHookResp struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func (l *feishuHookResp) Error() error {
	if l.Code == 0 {
		return nil
	}
	return merr.ErrorBadRequest("code: %d, msg: %s, data: %v", l.Code, l.Msg, l.Data)
}

func (f *feishuHook) Send(ctx context.Context, message Message) error {
	msg := make(map[string]any)
	if err := json.Unmarshal(message, &msg); err != nil {
		f.helper.Debugf("unmarshal feishu hook message failed: %v", err)
		return err
	}
	requestTime := time.Now().Unix()
	msg["timestamp"] = strconv.FormatInt(requestTime, 10)
	sign, err := f.sign(requestTime)
	if err != nil {
		f.helper.Debugf("gen feishu hook sign failed: %v", err)
		return err
	}
	msg["sign"] = sign
	requestBody, err := json.Marshal(msg)
	if err != nil {
		f.helper.Debugf("marshal feishu hook message failed: %v", err)
		return err
	}
	response, err := httpx.PostJson(ctx, f.api, requestBody)
	if err != nil {
		f.helper.Debugf("send feishu hook failed: %v", err)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		f.helper.Debugf("send feishu hook failed: status code: %d", response.StatusCode)
		return merr.ErrorBadRequest("status code: %d", response.StatusCode)
	}

	var resp feishuHookResp
	if err := json.NewDecoder(response.Body).Decode(&resp); err != nil {
		f.helper.Debugf("unmarshal feishu hook response failed: %v", err)
		return err
	}

	return resp.Error()
}

func (f *feishuHook) sign(timestamp int64) (string, error) {
	// timestamp + key sha256, then base64 encode
	signString := strconv.FormatInt(timestamp, 10) + "\n" + f.secret

	var data []byte
	h := hmac.New(sha256.New, []byte(signString))
	_, err := h.Write(data)
	if err != nil {
		f.helper.Debugf("gen feishu hook sign failed: %v", err)
		return "", err
	}

	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return signature, nil
}
