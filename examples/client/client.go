package client

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
	"errors"
	"net/http"
)

//counterfeiter:generate . MyHttpClient
type MyHttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type MyApplicationClient struct {
	HttpClient MyHttpClient
}

func NewApplicationClient(httpClient MyHttpClient) *MyApplicationClient {
	return &MyApplicationClient{
		HttpClient: httpClient,
	}
}

func (c *MyApplicationClient) PerformQuery() error {
	r, _ := http.NewRequest(http.MethodPost, "/resource", nil)
	resp, _ := c.HttpClient.Do(r)
	if resp.StatusCode == http.StatusUnauthorized {
		return errors.New("Unauthorized")
	}
	return nil
}
