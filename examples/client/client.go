package client

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

import (
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
