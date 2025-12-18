package zeptomail

type ZeptoMail struct {
	Email     Email
	FileCache FileCache
	Template  Template
}

// NewZeptoMail initializes the ZeptoMail client
func NewZeptoMail(mailAgent, apiKey, managementToken string) (*ZeptoMail, error) {
	emailClient, err := NewClient(mailAgent, apiKey)
	if err != nil {
		return nil, err
	}

	mgmtClient, err := NewClient(mailAgent, managementToken)
	if err != nil {
		return nil, err
	}

	return &ZeptoMail{
		Email:     Email(*emailClient),
		FileCache: FileCache(*emailClient),
		Template:  Template(*mgmtClient),
	}, nil
}
