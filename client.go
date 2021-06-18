package napv2

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strings"
)

type Client struct {
	Client *http.Client
	Auth   Authentication
}

func NewClient() *Client {
	return &Client{
		Client: http.DefaultClient,
	}
}

func (c *Client) SetAuth(auth Authentication) {
	c.Auth = auth
}

func buildRequest(method, url string, payload interface{}) (*http.Request, error) {
	if payload == nil {
		return http.NewRequest(method, url, nil)
	}
	b, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	buf := bytes.NewBuffer(b)
	return http.NewRequest(method, url, buf)
}

func (c *Client) ProcessRequest(baseURL string, res RestResource, params map[string]string, payload interface{}) error {
	trimmedEndpoint := strings.TrimLeft(res.RenderEndpoint(params), "/")
	trimmedBaseURL := strings.TrimRight(baseURL, "/")
	trimmedURL := trimmedBaseURL + "/" + trimmedEndpoint
	req, err := buildRequest(res.Method, trimmedURL, payload)
	if err != nil {
		return err
	}
	if c.Auth != nil {
		req.Header.Set("Authorization", c.Auth.AuthorizationHeader())
	}
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	return res.Router.CallFunc(resp)
}
