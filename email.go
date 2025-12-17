package zeptomail

import (
	"context"
	"net/http"
	"net/url"
)

type Email = Client

// SendHTMLEmail sends a HTML email
func (c *Email) SendHTMLEmail(ctx context.Context, req SendHTMLEmailReq) (*WrappedResponse[SendHTMLEmailRes], error) {
	path := "/email"
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[SendHTMLEmailReq, SendHTMLEmailRes](c, ctx, http.MethodPost, endpoint, nil, req)
}

// SendTemplatedEmail sends a templated email
func (c *Email) SendTemplatedEmail(ctx context.Context, req SendTemplatedEmailReq) (*WrappedResponse[SendTemplatedEmailRes], error) {
	path := "/email/template"
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[SendTemplatedEmailReq, SendTemplatedEmailRes](c, ctx, http.MethodPost, endpoint, nil, req)
}

// SendBatchTemplatedEmail sends a batch templated email
func (c *Email) SendBatchTemplatedEmail(ctx context.Context, req SendBatchTemplatedEmailReq) (*WrappedResponse[SendTemplatedEmailRes], error) {
	path := "/email/template/batch"
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[SendBatchTemplatedEmailReq, SendTemplatedEmailRes](c, ctx, http.MethodPost, endpoint, nil, req)
}

// SendBatchHTMLEmail The API is used to send a batch of transactional HTML emails.
func (c *Email) SendBatchHTMLEmail(ctx context.Context, req SendBatchHTMLEmailReq) (*WrappedResponse[SendBatchHTMLEmailRes], error) {
	path := "/email/batch"
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[SendBatchHTMLEmailReq, SendBatchHTMLEmailRes](c, ctx, http.MethodPost, endpoint, nil, req)
}
