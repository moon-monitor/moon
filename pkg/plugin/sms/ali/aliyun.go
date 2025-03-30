package ali

import (
	"context"
	"encoding/json"

	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapiV3 "github.com/alibabacloud-go/dysmsapi-20170525/v3/client"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/moon-monitor/moon/pkg/plugin/sms"

	"github.com/moon-monitor/moon/pkg/merr"
	"github.com/moon-monitor/moon/pkg/util/pointer"
)

var _ sms.Sender = (*aliyun)(nil)

const (
	endpoint = "dysmsapi.aliyuncs.com"
)

func NewAliyun(accessKeyID, accessKeySecret string, opts ...AliyunOption) (sms.Sender, error) {
	a := &aliyun{
		accessKeyID:     accessKeyID,
		accessKeySecret: accessKeySecret,
		endpoint:        endpoint,
	}
	for _, opt := range opts {
		opt(a)
	}
	if a.helper == nil {
		WithAliyunLogger(log.DefaultLogger)(a)
	}
	var err error
	a.clientV3, err = a.initV3()
	if err != nil {
		return nil, err
	}
	return a, nil
}

type AliyunOption func(*aliyun)

type aliyun struct {
	accessKeyID     string
	accessKeySecret string
	signName        string
	endpoint        string

	clientV3 *dysmsapiV3.Client
	helper   *log.Helper
}

// initV3 initializes the SMS clientV3
func (a *aliyun) initV3() (*dysmsapiV3.Client, error) {
	if a.accessKeySecret == "" || a.accessKeyID == "" {
		return nil, merr.ErrorBadRequest("SMS sending credential information is not configured")
	}
	config := &openapi.Config{
		AccessKeyId:     &a.accessKeyID,
		AccessKeySecret: &a.accessKeySecret,
		Endpoint:        tea.String(a.endpoint),
	}
	client, err := dysmsapiV3.NewClient(config)
	if err != nil {
		return nil, merr.ErrorBadRequest("Failed to initialize SMS clientV3").WithCause(err)
	}
	return client, nil
}

func (a *aliyun) Send(_ context.Context, message sms.Message) error {
	sendSmsRequest := &dysmsapiV3.SendSmsRequest{
		PhoneNumbers:  pointer.Of(message.PhoneNumber),
		SignName:      pointer.Of(a.signName),
		TemplateCode:  pointer.Of(message.Code),
		TemplateParam: pointer.Of(message.Content),
	}
	runtimeOptions := &util.RuntimeOptions{}
	response, err := a.clientV3.SendSmsWithOptions(sendSmsRequest, runtimeOptions)
	if err != nil {
		a.helper.Debugf("send sms failed: %v", err)
		return err
	}
	a.helper.Debugf("send sms response: %v", response)
	if pointer.Get(response.Body.Code) != "OK" {
		a.helper.Errorf("send sms failed: %v", response)
		body := pointer.Get(response.Body)
		return merr.ErrorBadRequest("send sms failed: %v", body)
	}
	return nil
}

func (a *aliyun) SendBatch(_ context.Context, messages []sms.Message) error {
	phoneNumbers := make([]string, 0, len(messages))
	signNames := make([]string, 0, len(messages))
	templateParams := make([]string, 0, len(messages))
	code := ""
	for _, message := range messages {
		phoneNumbers = append(phoneNumbers, message.PhoneNumber)
		signNames = append(signNames, a.signName)
		templateParams = append(templateParams, message.Content)
		code = message.Code
	}
	phoneNumberJson, err := json.Marshal(phoneNumbers)
	if err != nil {
		return merr.ErrorBadRequest("Failed to marshal phone numbers").WithCause(err)
	}
	signNameJson, err := json.Marshal(signNames)
	if err != nil {
		return merr.ErrorBadRequest("Failed to marshal sign names").WithCause(err)
	}
	templateParamJson, err := json.Marshal(templateParams)
	if err != nil {
		return merr.ErrorBadRequest("Failed to marshal template params").WithCause(err)
	}
	sendBatchSmsRequest := &dysmsapiV3.SendBatchSmsRequest{
		PhoneNumberJson:   pointer.Of(string(phoneNumberJson)),
		SignNameJson:      pointer.Of(string(signNameJson)),
		TemplateParamJson: pointer.Of(string(templateParamJson)),
		TemplateCode:      pointer.Of(code),
	}
	runtimeOptions := &util.RuntimeOptions{}
	response, err := a.clientV3.SendBatchSmsWithOptions(sendBatchSmsRequest, runtimeOptions)
	if err != nil {
		a.helper.Debugf("send batch sms failed: %v", err)
		return err
	}
	a.helper.Debugf("send batch sms response: %v", response)
	if pointer.Get(response.Body.Code) != "OK" {
		a.helper.Errorf("send batch sms failed: %v", response)
		body := pointer.Get(response.Body)
		return merr.ErrorBadRequest("send batch sms failed: %v", body)
	}
	return nil
}
