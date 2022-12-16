package version1

type EmailMessageV1 struct {
	From    string            `json:"from"`
	Cc      string            `json:"cc"`
	Bcc     string            `json:"bcc"`
	To      string            `json:"to"`
	ReplyTo string            `json:"reply_to"`
	Subject map[string]string `json:"subject"`
	Text    map[string]string `json:"text"`
	Html    map[string]string `json:"html"`
}

func EmptyEmailMessageV1() *EmailMessageV1 {
	return &EmailMessageV1{}
}

func NewEmailMessageV1(from string, to string, subject string, text string) *EmailMessageV1 {
	return &EmailMessageV1{
		From:    from,
		To:      to,
		Subject: map[string]string{"en": subject},
		Text:    map[string]string{"en": subject},
	}
}
