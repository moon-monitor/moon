package httpx

import (
	"context"
	"net/http"
	"net/url"
)

func NewClient() Client {
	return &config{
		Client: http.DefaultClient,
	}
}

type Client interface {
	Do(req *http.Request) (*http.Response, error)

	WithHeader(key, value string) Client
	WithBasicAuth(username, password string) Client
	WithContext(ctx context.Context) Client

	Get(api string, params ...url.Values) (*http.Response, error)
}

type config struct {
	*http.Client
}

type client struct {
	conf               *config
	header             http.Header
	username, password string
	ctx                context.Context
}

func (c *client) WithContext(ctx context.Context) Client {
	c.ctx = context.WithoutCancel(ctx)
	return c
}

func (c *client) Do(req *http.Request) (*http.Response, error) {
	if c.username != "" && c.password != "" {
		req.SetBasicAuth(c.username, c.password)
	}
	if c.ctx != nil {
		req = req.WithContext(c.ctx)
	}
	for k, h := range c.header {
		req.Header.Set(k, h[0])
	}
	return c.conf.Do(req)
}

func (c *client) WithHeader(key, value string) Client {
	c.header.Set(key, value)
	return c
}

func (c *client) WithBasicAuth(username, password string) Client {
	c.username = username
	c.password = password
	return c
}

func (c *client) Get(api string, params ...url.Values) (*http.Response, error) {
	urlParams := url.Values{}
	for _, param := range params {
		for k, v := range param {
			urlParams.Set(k, v[0])
		}
	}
	api = api + "?" + urlParams.Encode()
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}

func (c *config) WithHeader(key, value string) Client {
	cli := &client{
		conf:   c,
		header: make(http.Header),
		ctx:    context.Background(),
	}
	cli.header.Set(key, value)
	return cli
}

func (c *config) WithBasicAuth(username, password string) Client {
	return &client{
		conf:     c,
		header:   make(http.Header),
		username: username,
		password: password,
		ctx:      context.Background(),
	}
}

func (c *config) WithContext(ctx context.Context) Client {
	return &client{
		conf:     c,
		header:   make(http.Header),
		username: "",
		password: "",
		ctx:      ctx,
	}
}

func (c *config) Get(api string, params ...url.Values) (*http.Response, error) {
	urlParams := url.Values{}
	for _, param := range params {
		for k, v := range param {
			urlParams.Set(k, v[0])
		}
	}
	api = api + "?" + urlParams.Encode()
	req, err := http.NewRequest(http.MethodGet, api, nil)
	if err != nil {
		return nil, err
	}
	return c.Do(req)
}
