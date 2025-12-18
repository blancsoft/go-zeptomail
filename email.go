package zeptomail

import (
	"context"
	"net/http"
)

type Email Client

// SendHTMLEmail sends a HTML email
func (e *Email) SendHTMLEmail(ctx context.Context, req SendHTMLEmailReq) (*WrappedResponse[SendHTMLEmailRes], error) {
	path := "/email"
	endpoint := e.baseURL.JoinPath(path)
	return request[SendHTMLEmailReq, SendHTMLEmailRes]((*Client)(e), ctx, http.MethodPost, endpoint, nil, req)
}

// SendBatchHTMLEmail The API is used to send a batch of transactional HTML emails.
func (e *Email) SendBatchHTMLEmail(ctx context.Context, req SendBatchHTMLEmailReq) (*WrappedResponse[SendBatchHTMLEmailRes], error) {
	path := "/email/batch"
	endpoint := e.baseURL.JoinPath(path)
	return request[SendBatchHTMLEmailReq, SendBatchHTMLEmailRes]((*Client)(e), ctx, http.MethodPost, endpoint, nil, req)
}

// SendTemplatedEmail sends a templated email
func (e *Email) SendTemplatedEmail(ctx context.Context, req SendTemplatedEmailReq) (*WrappedResponse[SendTemplatedEmailRes], error) {
	path := "/email/template"
	endpoint := e.baseURL.JoinPath(path)
	return request[SendTemplatedEmailReq, SendTemplatedEmailRes]((*Client)(e), ctx, http.MethodPost, endpoint, nil, req)
}

// SendBatchTemplatedEmail sends a batch templated email
func (e *Email) SendBatchTemplatedEmail(ctx context.Context, req SendBatchTemplatedEmailReq) (*WrappedResponse[SendTemplatedEmailRes], error) {
	path := "/email/template/batch"
	endpoint := e.baseURL.JoinPath(path)
	return request[SendBatchTemplatedEmailReq, SendTemplatedEmailRes]((*Client)(e), ctx, http.MethodPost, endpoint, nil, req)
}
