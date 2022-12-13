package test_version1

import (
	"context"
	"testing"

	"github.com/pip-services-infrastructure2/client-email-go/version1"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/stretchr/testify/assert"
)

type EmailClientFixtureV1 struct {
	Client version1.IEmailClientV1
}

func NewEmailClientFixtureV1(client version1.IEmailClientV1) *EmailClientFixtureV1 {
	return &EmailClientFixtureV1{
		Client: client,
	}
}

func (c *EmailClientFixtureV1) TestSendEmailToAddress(t *testing.T) {
	message := &version1.EmailMessageV1{
		To:      "somebody@somewhere.com",
		Subject: map[string]string{"en": "{{subject}}"},
		Text:    map[string]string{"en": "{{text}}"},
		Html:    map[string]string{"en": "<p>{{text}}</p>"},
	}

	parameters := config.NewConfigParamsFromTuples(
		"subject", "Test Email to Address",
		"text", "This is just a test",
	)

	err := c.Client.SendMessage(context.Background(), "", message, parameters)
	assert.Nil(t, err)
}

func (c *EmailClientFixtureV1) TestSendEmailToRecipient(t *testing.T) {
	message := &version1.EmailMessageV1{
		Subject: map[string]string{"en": "Test Email to Recipient"},
		Text:    map[string]string{"en": "This is just a test email"},
	}

	recipient := &version1.EmailRecipientV1{
		Id:    "1",
		Email: "somebody@somewhere.com",
		Name:  "Test Recipient",
	}

	err := c.Client.SendMessageToRecipient(context.Background(), "", recipient, message, nil)
	assert.Nil(t, err)
}
