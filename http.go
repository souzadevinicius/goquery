package goquery

import (
	"net/http"
)

type Http interface {
	Get(url string) (resp *http.Response, err error)
}

type HttpClient struct {
	http Http
}

func New(http Http) *HttpClient {
	return &HttpClient{http}
}

func (c *HttpClient) Get(url string) (resp *http.Response, err error) {
	return c.http.Get(url)
}
