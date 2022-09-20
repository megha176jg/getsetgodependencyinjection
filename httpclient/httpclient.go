package httpclient

import (
	"net/http"
	"time"
)

func New(timeout int) http.Client {
	t := http.DefaultTransport.(*http.Transport).Clone()
	t.MaxIdleConns = 100
	t.MaxConnsPerHost = 100
	t.MaxIdleConnsPerHost = 100
	timeoutDur := time.Duration(timeout) * time.Second
	httpClient := http.Client{
		Timeout:   timeoutDur,
		Transport: t,
	}
	return httpClient
}

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type httpClient struct {
	client http.Client
}

func NewHttpClient(timeout int) HTTPClient {
	return &httpClient{client: New(timeout)}
}

func (c *httpClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
