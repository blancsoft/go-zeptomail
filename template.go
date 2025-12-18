package zeptomail

import (
	"context"
	"fmt"
	"net/http"
)

type Template Client

// AddEmailTemplate is used to add an email template.
func (t *Template) AddEmailTemplate(ctx context.Context, req AddEmailTemplateReq) (*WrappedResponse[AddEmailTemplateRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates", t.mailAgent)
	endpoint := t.baseURL.JoinPath(path)
	return request[AddEmailTemplateReq, AddEmailTemplateRes]((*Client)(t), ctx, http.MethodPost, endpoint, nil, req)
}

// UpdateEmailTemplate is used to update an email template.
func (t *Template) UpdateEmailTemplate(ctx context.Context, req UpdateEmailTemplateReq) (*WrappedResponse[AddEmailTemplateRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates/%s", t.mailAgent, req.TemplateKey)
	endpoint := t.baseURL.JoinPath(path)
	return request[UpdateEmailTemplateReq, AddEmailTemplateRes]((*Client)(t), ctx, http.MethodPut, endpoint, nil, req)
}

// ListEmailTemplates lists the required number of email templates in your ZeptoMail account.
func (t *Template) ListEmailTemplates(ctx context.Context, Offset, limit int) (*WrappedResponse[ListEmailTemplatesRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates?offset=%d&limit=%d", t.mailAgent, Offset, limit)
	endpoint := t.baseURL.JoinPath(path)
	return request[any, ListEmailTemplatesRes]((*Client)(t), ctx, http.MethodGet, endpoint, nil, nil)
}

// GetEmailTemplate is used to fetch a particular email template.
func (t *Template) GetEmailTemplate(ctx context.Context, TemplateKey string) (*WrappedResponse[GetEmailTemplateRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates/%s", t.mailAgent, TemplateKey)
	endpoint := t.baseURL.JoinPath(path)
	return request[any, GetEmailTemplateRes]((*Client)(t), ctx, http.MethodGet, endpoint, nil, nil)
}

// DeleteEmailTemplate is used to delete a template using template key.
func (t *Template) DeleteEmailTemplate(ctx context.Context, TemplateKey string) (*WrappedResponse[any], error) {
	path := fmt.Sprintf("/mailagents/%s/templates/%s", t.mailAgent, TemplateKey)
	endpoint := t.baseURL.JoinPath(path)
	return request[any, any]((*Client)(t), ctx, http.MethodDelete, endpoint, nil, nil)
}
