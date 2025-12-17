package zeptomail

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"

	"github.com/go-playground/validator/v10"
)

type Client struct {
	client  *http.Client
	baseURL *url.URL
	apiKey  string
}

func NewClient(baseUrl, apiKey string, defaultClient ...*http.Client) (*Client, error) {
	u, err := url.Parse(baseUrl)
	if err != nil {
		return nil, err
	}
	if u.Scheme == "" {
		return nil, fmt.Errorf("url scheme is required")
	}

	var httpClient *http.Client
	if len(defaultClient) > 0 && defaultClient[0] != nil {
		httpClient = defaultClient[0]
	} else {
		httpClient = &http.Client{
			Transport: http.DefaultTransport.(*http.Transport).Clone(),
		}
	}

	return &Client{
		client:  httpClient,
		baseURL: u,
		apiKey:  apiKey,
	}, nil
}

type WrappedResponse[T any] struct {
	RawResponse *http.Response
	Data        T
}

// validate runs a validation on the incoming json payload
var validate = validator.New(validator.WithRequiredStructEnabled())

func request[S any, R any](
	c *Client, ctx context.Context,
	method string, endpoint *url.URL,
	headers http.Header, payload S,
) (*WrappedResponse[R], error) {
	if err := validate.Struct(&payload); err != nil && payload != nil {
		return nil, err
	}

	var buff bytes.Buffer
	if err := json.NewEncoder(&buff).Encode(payload); err != nil && payload != nil {
		return nil, fmt.Errorf("encoding failed: %w", err)
	}

	req, err := http.NewRequest(method, endpoint.String(), &buff)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %w", err)
	}

	req = req.WithContext(ctx)
	if buff.Len() != 0 {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Authorization", c.apiKey)
	for k, v := range headers {
		req.Header[k] = v
	}

	var rv WrappedResponse[R]
	rv.RawResponse, err = c.client.Do(req)
	if err != nil {
		return &rv, fmt.Errorf("request failed: %w", err)
	}
	rv.RawResponse.Body = io.NopCloser(rv.RawResponse.Body)

	if err = json.NewDecoder(rv.RawResponse.Body).Decode(&rv.Data); err != nil {
		return &rv, fmt.Errorf("decoding failed: %w", err)
	}
	return &rv, nil
}
