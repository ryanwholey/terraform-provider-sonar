package sonar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

// Client struct
type Client struct {
	client    *retryablehttp.Client
	BaseURL   *url.URL
	AuthToken string
}

// DoRequest Make request on behalf of client
func (c *Client) DoRequest(method, path string, body, v interface{}) (*http.Response, error) {
	parsedURL, err := url.Parse(path)
	if err != nil {
		return nil, err
	}

	resolvedURL := c.BaseURL.ResolveReference(parsedURL)

	buffer := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buffer).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := retryablehttp.NewRequest(method, resolvedURL.String(), buffer)
	if err != nil {
		return nil, err
	}

	req.Close = true

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.AuthToken))
	req.Header.Add("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, err
}

// PacketRetryPolicy to satisfy the retry library
func PacketRetryPolicy(ctx context.Context, res *http.Response, err error) (bool, error) {
	return false, nil
}

// NewClient creates a new client
func NewClient(authToken string, baseURL string) (*Client, error) {
	httpClient := retryablehttp.NewClient()
	httpClient.RetryWaitMin = time.Second
	httpClient.RetryWaitMax = 30 * time.Second
	httpClient.RetryMax = 10
	httpClient.CheckRetry = PacketRetryPolicy

	u, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	c := &Client{client: httpClient, BaseURL: u, AuthToken: authToken}

	return c, nil
}
