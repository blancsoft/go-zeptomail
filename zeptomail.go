package zeptomail

type ZeptoMail struct {
	Email     Email
	Template  Template
	FileCache FileCache
}

// NewZeptoMail initializes the ZeptoMail client
func NewZeptoMail(baseUrl, apiKey string) (*ZeptoMail, error) {
	client, err := NewClient(baseUrl, apiKey)
	if err != nil {
		return nil, err
	}

	return &ZeptoMail{
		Email:     *client,
		Template:  *client,
		FileCache: *client,
	}, nil
}
