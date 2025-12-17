package zeptomail

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
)

type Template = Client

// AddEmailTemplate is used to add an email template.
func (c *Template) AddEmailTemplate(ctx context.Context, req AddEmailTemplateReq) (*WrappedResponse[AddEmailTemplateRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates", req.MailagentAlias)
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[AddEmailTemplateReq, AddEmailTemplateRes](c, ctx, http.MethodPost, endpoint, nil, req)
}

// UpdateEmailTemplate is used to update an email template.
func (c *Template) UpdateEmailTemplate(ctx context.Context, req UpdateEmailTemplateReq) (*WrappedResponse[AddEmailTemplateRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates/%s", req.MailagentAlias, req.TemplateKey)
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[UpdateEmailTemplateReq, AddEmailTemplateRes](c, ctx, http.MethodPut, endpoint, nil, req)
}

// ListEmailTemplates lists the required number of email templates in your ZeptoMail account.
func (c *Template) ListEmailTemplates(ctx context.Context, MailAgentAlias string, Offset, limit int) (*WrappedResponse[ListEmailTemplatesRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates?offset=%d&limit=%d", MailAgentAlias, Offset, limit)
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[any, ListEmailTemplatesRes](c, ctx, http.MethodGet, endpoint, nil, nil)
}

// GetEmailTemplate is used to fetch a particular email template.
func (c *Template) GetEmailTemplate(ctx context.Context, MailAgentAlias, TemplateKey string) (*WrappedResponse[GetEmailTemplateRes], error) {
	path := fmt.Sprintf("/mailagents/%s/templates/%s", MailAgentAlias, TemplateKey)
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[any, GetEmailTemplateRes](c, ctx, http.MethodGet, endpoint, nil, nil)
}

// DeleteEmailTemplate is used to delete a template using template key.
func (c *Template) DeleteEmailTemplate(ctx context.Context, MailAgentAlias, TemplateKey string) (*WrappedResponse[any], error) {
	path := fmt.Sprintf("/mailagents/%s/templates/%s", MailAgentAlias, TemplateKey)
	endpoint := c.baseURL.ResolveReference(&url.URL{Path: path})
	return request[any, any](c, ctx, http.MethodDelete, endpoint, nil, nil)
}
