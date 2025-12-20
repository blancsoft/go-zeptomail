package zeptomail

import (
	"fmt"
	"strings"
)

type ZeptoMail struct {
	Email     Email
	FileCache FileCache
	Template  Template
}

// NewZeptoMail initializes the ZeptoMail client
func NewZeptoMail(mailAgent, apiKey, oauthToken string) (*ZeptoMail, error) {
	const (
		apiKeyPrefix     = "Zoho-enczapikey"
		oauthTokenPrefix = "Zoho-oauthtoken"
	)
	if apiKey != "" && !strings.HasPrefix(apiKey, apiKeyPrefix) {
		apiKey = fmt.Sprintf("%s %s", apiKeyPrefix, strings.TrimSpace(apiKey))
	}

	if oauthToken != "" && !strings.HasPrefix(oauthToken, "Zoho-oauthtoken") {
		oauthToken = fmt.Sprintf("%s %s", oauthTokenPrefix, strings.TrimSpace(oauthToken))
	}

	emailClient, err := NewClient(mailAgent, apiKey)
	if err != nil {
		return nil, err
	}

	mgmtClient, err := NewClient(mailAgent, oauthToken)
	if err != nil {
		return nil, err
	}

	return &ZeptoMail{
		Email:     Email(*emailClient),
		FileCache: FileCache(*emailClient),
		Template:  Template(*mgmtClient),
	}, nil
}
