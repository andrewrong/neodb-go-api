package util

import (
	"bytes"
	"net/http"
	"time"
)

type HttpClient struct {
	client *http.Client
	auth   string
}

func NewHttpClient(auth string, timeoutMS int, tr http.RoundTripper) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Timeout:   time.Duration(timeoutMS) * time.Millisecond,
			Transport: tr,
		},
		auth: auth,
	}
}
func (c *HttpClient) Get(url string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.auth)
	return c.client.Do(req)
}
func (c *HttpClient) Post(url string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.auth)
	req.Header.Add("Content-Type", "application/json")
	return c.client.Do(req)
}
func (c *HttpClient) Delete(url string) (*http.Response, error) {
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", c.auth)
	return c.client.Do(req)
}
