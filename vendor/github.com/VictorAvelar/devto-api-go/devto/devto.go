package devto

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/url"
)

//Configuration constants
const (
	BaseURL      string = "https://dev.to"
	APIVersion   string = "0.5.1"
	APIKeyHeader string = "api-key"
)

//devto client errors
var (
	ErrMissingConfig     = errors.New("missing configuration")
	ErrProtectedEndpoint = errors.New("to use this resource you need to provide an authentication method")
)

type httpClient interface {
	Do(req *http.Request) (res *http.Response, err error)
}

//Client is the main data structure for performing actions
//against dev.to API
type Client struct {
	Context    context.Context
	BaseURL    *url.URL
	HTTPClient httpClient
	Config     *Config
	Articles   *ArticlesResource
}

//NewClient takes a context, a configuration pointer and optionally a
//base http client (bc) to build an Client instance.
func NewClient(ctx context.Context, conf *Config, bc httpClient, bu string) (dev *Client, err error) {
	if bc == nil {
		bc = http.DefaultClient
	}

	if ctx == nil {
		ctx = context.Background()
	}

	if conf == nil {
		return nil, ErrMissingConfig
	}

	if bu == "" {
		bu = BaseURL
	}

	u, _ := url.Parse(bu)

	c := &Client{
		Context:    ctx,
		BaseURL:    u,
		HTTPClient: bc,
		Config:     conf,
	}
	c.Articles = &ArticlesResource{API: c}
	return c, nil
}

//NewRequest build the request relative to the client BaseURL
func (c *Client) NewRequest(method string, uri string, body io.Reader) (*http.Request, error) {
	u, err := url.Parse(uri)
	if err != nil {
		return nil, err
	}
	fu := c.BaseURL.ResolveReference(u).String()
	req, err := http.NewRequest(method, fu, body)
	if err != nil {
		return nil, err
	}
	return req, nil
}
