package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type EmailCommandableHttpClientV1 struct {
	*clients.CommandableHttpClient
}

func NewEmailCommandableHttpClientV1() *EmailCommandableHttpClientV1 {
	return NewEmailCommandableHttpClientV1WithConfig(nil)
}

func NewEmailCommandableHttpClientV1WithConfig(config *cconf.ConfigParams) *EmailCommandableHttpClientV1 {
	c := &EmailCommandableHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/email"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *EmailCommandableHttpClientV1) SendMessage(ctx context.Context, correlationId string, message *EmailMessageV1, parameters *cconf.ConfigParams) error {
	params := cdata.NewAnyValueMapFromTuples(
		"message", message,
		"parameters", parameters,
	)

	_, err := c.CallCommand(ctx, "send_message", correlationId, params)

	return err
}

func (c *EmailCommandableHttpClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipient *EmailRecipientV1, message *EmailMessageV1, parameters *cconf.ConfigParams) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient", recipient,
		"message", message,
		"parameters", parameters,
	)

	_, err := c.CallCommand(ctx, "send_message_to_recipient", correlationId, params)

	return err
}

func (c *EmailCommandableHttpClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*EmailRecipientV1, message *EmailMessageV1, parameters *cconf.ConfigParams) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipients", recipients,
		"message", message,
		"parameters", parameters,
	)

	_, err := c.CallCommand(ctx, "send_message_to_recipients", correlationId, params)

	return err
}
