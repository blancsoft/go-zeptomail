package zeptomail_test

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/blancsoft/go-zeptomail"
)

var (
	attachment []zeptomail.EmailAttachment
	sender     = zeptomail.EmailAddress{
		Address: "emailtesting.sender@blancsoft.com",
		Name:    "Blancsoft Tester",
	}
	receiver = zeptomail.EmailAddress{
		Address: "emailtesting.receiver@blancsoft.com",
		Name:    "Blancsoft Receiver",
	}
	other = zeptomail.EmailAddress{
		Address: "emailtesting.other@blancsoft.com",
		Name:    "Blancsoft Other",
	}
	testHeaders = map[string]any{
		"X-Tester": "go-zeptomail",
	}
	emailSubject = "{{testType}} — Invitation to the event"
	emailBody    = `<h1>Hello {{name}},</h1><p>Your invitation link  <a href="{{link}}">{{link}}</a></p>`
)

func init() {
	b64Attachment := make([]byte, base64.StdEncoding.EncodedLen(len(fileAttachment)))
	base64.StdEncoding.Encode(b64Attachment, fileAttachment)
	attachment = []zeptomail.EmailAttachment{
		{Content: string(b64Attachment), Name: "attachment.ico", MimeType: "image/x-icon"},
	}
}

func TestHTMLEmail(t *testing.T) {
	zepto, err := zeptomail.NewZeptoMail(zeptoMailAgent, zeptoMailToken, zeptoMailMgmtToken)
	require.NoError(t, err)

	t.Run("send html email", func(t *testing.T) {
		ref := strings.ToLower(rand.Text()[:10])
		rv, err := zepto.Email.SendHTMLEmail(t.Context(), zeptomail.SendHTMLEmailReq{
			BaseSendEmail: zeptomail.BaseSendEmail{
				From:      sender,
				To:        []zeptomail.SendEmailTo{{EmailAddress: receiver}},
				MergeInfo: map[string]any{"name": "World", "link": "https://blancsoft.com", "testType": "HTML"},
			},
			BaseEmailOption: zeptomail.BaseEmailOption{
				CC:              []zeptomail.SendEmailTo{{EmailAddress: receiver}},
				BCC:             []zeptomail.SendEmailTo{{EmailAddress: receiver}},
				ReplyTo:         []zeptomail.EmailAddress{sender},
				TrackClicks:     true,
				TrackOpens:      true,
				ClientReference: ref,
				MimeHeaders:     testHeaders,
				Attachments:     attachment,
			},
			Subject:  emailSubject,
			HtmlBody: emailBody,
		})
		{ // TODO: Remove me
			bb, err := io.ReadAll(rv.RawResponse.Body)
			require.NoError(t, err)
			fmt.Printf("RAW RESPONSE: %s\n", string(bb))
		}
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)
			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, "email", rv.Data.Object)
			assert.NotEmpty(t, rv.Data.RequestId)
			assert.NotNil(t, rv.Data.Data)
			assert.Equal(t, "Email request received", rv.Data.Data[0].Message)
		})
	})

	t.Run("send batch html email", func(t *testing.T) {
		rv, err := zepto.Email.SendBatchHTMLEmail(t.Context(), zeptomail.SendBatchHTMLEmailReq{
			From: sender,
			To: []zeptomail.SendBatchEmailTo{
				{EmailAddress: receiver, MergeInfo: map[string]any{"name": "Other", "link": "https://blancsoft.com", "testType": "Batch HTML"}},
				{EmailAddress: other, MergeInfo: map[string]any{"name": "Sender", "link": "https://blancsoft.com", "testType": "Batch HTML"}},
			},
			Subject:  emailSubject,
			HtmlBody: emailBody,
		})
		{ // TODO: Remove me
			bb, err := io.ReadAll(rv.RawResponse.Body)
			require.NoError(t, err)
			fmt.Printf("RAW RESPONSE: %s\n", string(bb))
		}
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)
			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, "email", rv.Data.Object)
			assert.NotEmpty(t, rv.Data.RequestId)
			assert.NotNil(t, rv.Data.Data)
			assert.Equal(t, "Email request received", rv.Data.Data[0].Message)
		})
	})
}

func setupTemplated(t *testing.T, zepto *zeptomail.ZeptoMail, ref string) zeptomail.AddEmailTemplateRes {
	t.Helper()

	tmpl := zeptomail.AddEmailTemplateReq{
		TemplateName:  "E-invite",
		TemplateAlias: "en_invite_" + ref,
		Subject:       emailSubject,
		HtmlBody:      emailBody,
	}
	addTmplRv, err := zepto.Template.AddEmailTemplate(t.Context(), tmpl)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, addTmplRv.RawResponse.StatusCode)
	t.Cleanup(func() {
		fmt.Printf("TODO: cleaning up\n")
		delTmplRv, err := zepto.Template.DeleteEmailTemplate(context.Background(), addTmplRv.Data.Data.TemplateKey)
		require.NoError(t, err)
		require.Equal(t, http.StatusNoContent, delTmplRv.RawResponse.StatusCode)
	})

	return addTmplRv.Data
}

func TestTemplateEmail(t *testing.T) {
	zepto, err := zeptomail.NewZeptoMail(zeptoMailAgent, zeptoMailToken, zeptoMailMgmtToken)
	require.NoError(t, err)

	t.Run("send templated email", func(t *testing.T) {
		ref := strings.ToLower(rand.Text()[:10])
		tmplRv := setupTemplated(t, zepto, ref)
		rv, err := zepto.Email.SendTemplatedEmail(t.Context(), zeptomail.SendTemplatedEmailReq{
			TemplateKey: tmplRv.Data.TemplateKey,
			BaseSendEmail: zeptomail.BaseSendEmail{
				From:      sender,
				To:        []zeptomail.SendEmailTo{{EmailAddress: receiver}},
				MergeInfo: map[string]any{"name": "World", "link": "https://blancsoft.com", "testType": "Templated"},
			},
			BaseEmailOption: zeptomail.BaseEmailOption{
				CC:              []zeptomail.SendEmailTo{{EmailAddress: receiver}},
				BCC:             []zeptomail.SendEmailTo{{EmailAddress: receiver}},
				ReplyTo:         []zeptomail.EmailAddress{sender},
				TrackClicks:     true,
				TrackOpens:      true,
				ClientReference: ref,
				MimeHeaders:     testHeaders,
				Attachments:     attachment,
			},
		})
		{ // TODO: Remove me
			bb, err := io.ReadAll(rv.RawResponse.Body)
			require.NoError(t, err)
			fmt.Printf("RAW RESPONSE: %s\n", string(bb))
		}
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)
			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, "email", rv.Data.Object)
			assert.NotEmpty(t, rv.Data.RequestId)
			assert.NotNil(t, rv.Data.Data)
			assert.Equal(t, "Email request received", rv.Data.Data[0].Message)
		})
	})

	t.Run("send batch templated email", func(t *testing.T) {
		ref := strings.ToLower(rand.Text()[:10])
		tmplRv := setupTemplated(t, zepto, ref)
		rv, err := zepto.Email.SendBatchTemplatedEmail(t.Context(), zeptomail.SendBatchTemplatedEmailReq{
			TemplateKey:   tmplRv.Data.TemplateKey,
			BounceAddress: "",
			From:          sender,
			To: []zeptomail.SendBatchEmailTo{
				{EmailAddress: receiver, MergeInfo: map[string]any{"name": "Receiver", "link": "https://blancsoft.com", "testType": "Batch Templated"}},
				{EmailAddress: other, MergeInfo: map[string]any{"name": "Other", "link": "https://blancsoft.com", "testType": "Batch Templated"}},
			},
			ReplyTo:         sender,
			TrackClicks:     true,
			TrackOpens:      true,
			ClientReference: ref,
			MimeHeaders:     testHeaders,
			Attachments:     attachment,
		})
		{ // TODO: Remove me
			bb, err := io.ReadAll(rv.RawResponse.Body)
			require.NoError(t, err)
			fmt.Printf("RAW RESPONSE: %s\n", string(bb))
		}
		require.NoError(t, err)

		t.Run("status code", func(t *testing.T) {
			assert.Equal(t, http.StatusCreated, rv.RawResponse.StatusCode)
		})

		t.Run("valid response", func(t *testing.T) {
			assert.Nil(t, rv.Data.Error)
			assert.Equal(t, "OK", rv.Data.Message)
			assert.Equal(t, "email", rv.Data.Object)
			assert.NotEmpty(t, rv.Data.RequestId)
			assert.NotNil(t, rv.Data.Data)
			assert.Equal(t, "Email request received", rv.Data.Data[0].Message)
		})
	})
}
