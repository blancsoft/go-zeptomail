package zeptomail_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/blancsoft/go-zeptomail"
)

func TestZeptoMailTemplate(t *testing.T) {
	//t.Skip("Requires Oauth2 Token. Complicated process")
	tmpl := zeptomail.AddEmailTemplateReq{
		TemplateName:  "E-invite",
		Subject:       "Invitation to the event",
		HtmlBody:      "<h1> Hi {{name}}</h1>, invitation link {{link}}",
		TemplateAlias: "en_invite",
	}
	tmplUpdate := zeptomail.UpdateEmailTemplateReq{
		TemplateName: "Invite Link",
		Subject:      "Event Invitation",
		TextBody:     "Hello Guest, your invitation link is {{link}}",
	}

	zepto, err := zeptomail.NewZeptoMail(zeptoMailAgent, "", zeptoMailMgmtToken)
	require.NoError(t, err)

	t.Run("add template", func(t *testing.T) {
		rv, err := zepto.Template.AddEmailTemplate(t.Context(), tmpl)
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)

			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, tmpl.TemplateName, rv.Data.Data.TemplateName)
			assert.Equal(t, tmpl.Subject, rv.Data.Data.Subject)
			assert.Contains(t, rv.Data.Data.HtmlBody, `<h1>Hi {{name}}</h1>`)
			assert.NotEmpty(t, rv.Data.Data.TemplateKey)
			assert.NotEmpty(t, rv.Data.Data.SampleMergeInfo)
			assert.NotEmpty(t, rv.Data.Data.CreatedTime)
			assert.NotEmpty(t, rv.Data.Data.ModifiedTime)

			// set an update template key
			tmplUpdate.TemplateKey = rv.Data.Data.TemplateKey
		})
	})

	t.Run("list templates", func(t *testing.T) {
		t.Skip("unexpected error occurs: `This URL does not exist`")
		rv, err := zepto.Template.ListEmailTemplates(t.Context(), 0, 10)
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)

			assert.Equal(t, "OK", rv.Data.Message)
			assert.NotNil(t, rv.Data.Data)
			assert.NotEmpty(t, rv.Data.Data[0].TemplateName)
			assert.NotEmpty(t, rv.Data.Data[0].Subject)
			assert.NotEmpty(t, rv.Data.Data[0].TemplateKey)
			assert.NotEmpty(t, rv.Data.Data[0].TemplateAlias)
			assert.NotEmpty(t, rv.Data.Data[0].CreatedTime)
			assert.NotEmpty(t, rv.Data.Data[0].ModifiedTime)
		})
	})

	t.Run("update template", func(t *testing.T) {
		rv, err := zepto.Template.UpdateEmailTemplate(t.Context(), tmplUpdate)
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusOK, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)

			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, tmplUpdate.Subject, rv.Data.Data.Subject)
			assert.Equal(t, tmplUpdate.HtmlBody, rv.Data.Data.HtmlBody)
			assert.Equal(t, tmplUpdate.TextBody, rv.Data.Data.TextBody)
			assert.Equal(t, tmpl.TemplateAlias, rv.Data.Data.TemplateAlias)
			assert.Equal(t, tmplUpdate.TemplateName, rv.Data.Data.TemplateName)
			assert.NotEmpty(t, rv.Data.Data.TemplateKey)
			assert.NotEmpty(t, rv.Data.Data.TemplateLink)
			assert.NotEmpty(t, rv.Data.Data.SampleMergeInfo)
			assert.NotEmpty(t, rv.Data.Data.CreatedTime)
			assert.NotEmpty(t, rv.Data.Data.ModifiedTime)
		})
	})

	t.Run("get template", func(t *testing.T) {
		rv, err := zepto.Template.GetEmailTemplate(t.Context(), tmplUpdate.TemplateKey)
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)

			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, "templates", rv.Data.Object)
			assert.Equal(t, tmplUpdate.Subject, rv.Data.Data.Subject)
			assert.Equal(t, tmplUpdate.HtmlBody, rv.Data.Data.HtmlBody)
			assert.Equal(t, tmplUpdate.TextBody, rv.Data.Data.TextBody)
			assert.Equal(t, tmpl.TemplateAlias, rv.Data.Data.TemplateAlias)
			assert.Equal(t, tmplUpdate.TemplateName, rv.Data.Data.TemplateName)
			assert.NotEmpty(t, rv.Data.Data.TemplateKey)
			assert.NotEmpty(t, rv.Data.Data.TemplateLink)
			assert.NotEmpty(t, rv.Data.Data.SampleMergeInfo)
			assert.NotEmpty(t, rv.Data.Data.CreatedTime)
			assert.NotEmpty(t, rv.Data.Data.ModifiedTime)
		})
	})

	t.Run("delete template", func(t *testing.T) {
		rv, err := zepto.Template.DeleteEmailTemplate(t.Context(), tmplUpdate.TemplateKey)
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusNoContent, rv.RawResponse.StatusCode)
		})
	})
}
