package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/cs2cap/cli/internal/models"
)

const DefaultBaseURL = "https://api.cs2c.app"
const DefaultTimeout = 30 * time.Second

type Client struct {
	baseURL    string
	apiKey     string
	httpClient *http.Client
}

func NewClient(apiKey, baseURL string) *Client {
	if baseURL == "" {
		baseURL = DefaultBaseURL
	}
	return &Client{
		baseURL: baseURL,
		apiKey:  apiKey,
		httpClient: &http.Client{
			Timeout: DefaultTimeout,
		},
	}
}

func (c *Client) get(ctx context.Context, path string, query url.Values, v interface{}) error {
	u, err := url.Parse(c.baseURL + path)
	if err != nil {
		return fmt.Errorf("invalid url: %w", err)
	}
	if query != nil {
		u.RawQuery = query.Encode()
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	c.setAuth(req)

	return c.do(req, v)
}

func (c *Client) post(ctx context.Context, path string, body, v interface{}) error {
	var buf bytes.Buffer
	if body != nil {
		if err := json.NewEncoder(&buf).Encode(body); err != nil {
			return fmt.Errorf("encode body: %w", err)
		}
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, c.baseURL+path, &buf)
	if err != nil {
		return fmt.Errorf("create request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	c.setAuth(req)

	return c.do(req, v)
}

func (c *Client) setAuth(req *http.Request) {
	if c.apiKey != "" {
		req.Header.Set("Authorization", "Bearer "+c.apiKey)
	}
}

func (c *Client) do(req *http.Request, v interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("read response: %w", err)
	}

	if resp.StatusCode >= 400 {
		var errResp models.ErrorResponse
		if err := json.Unmarshal(body, &errResp); err == nil && errResp.Code != "" {
			return fmt.Errorf("API error [%d]: %s: %s", resp.StatusCode, errResp.Code, errResp.Detail)
		}
		return fmt.Errorf("API error [%d]: %s", resp.StatusCode, string(body))
	}

	if v != nil {
		if err := json.Unmarshal(body, v); err != nil {
			return fmt.Errorf("decode response: %w", err)
		}
	}

	return nil
}

func joinInts(items []int) []string {
	s := make([]string, len(items))
	for i, v := range items {
		s[i] = fmt.Sprintf("%d", v)
	}
	return s
}
