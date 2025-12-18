package zeptomail

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"

	"github.com/go-playground/validator/v10"
)

const baseURL = "https://api.zeptomail.com/v1.1"

type Client struct {
	client    *http.Client
	baseURL   *url.URL
	mailAgent string
	apiKey    string
}

func NewClient(mailAgent, apiKey string, defaultClient ...*http.Client) (*Client, error) {
	u, err := url.Parse(baseURL)
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
		client:    httpClient,
		baseURL:   u,
		mailAgent: mailAgent,
		apiKey:    apiKey,
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
	v := reflect.ValueOf(payload)
	hasPayload := v.IsValid() && !v.IsZero()

	var buff bytes.Buffer
	if hasPayload {
		if err := validate.Struct(&payload); err != nil {
			return nil, err
		}

		if err := json.NewEncoder(&buff).Encode(payload); err != nil {
			return nil, fmt.Errorf("encoding failed: %w", err)
		}
	}

	req, err := http.NewRequest(method, endpoint.String(), &buff)
	if err != nil {
		return nil, fmt.Errorf("new request failed: %w", err)
	}

	req = req.WithContext(ctx)
	if hasPayload {
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

	var body bytes.Buffer
	rv.RawResponse.Body = io.NopCloser(io.TeeReader(rv.RawResponse.Body, &body))

	if err = json.NewDecoder(rv.RawResponse.Body).Decode(&rv.Data); err != nil && !errors.Is(err, io.EOF) {
		return &rv, fmt.Errorf("decoding failed: %w", err)
	}
	rv.RawResponse.Body = io.NopCloser(&body)
	return &rv, nil
}
