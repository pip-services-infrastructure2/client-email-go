package version1

import (
	"context"

	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cdata "github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-rpc-gox/clients"
)

type EmailHttpClientV1 struct {
	*clients.CommandableHttpClient
	defaultParameters *cconf.ConfigParams
}

func NewEmailHttpClientV1() *EmailHttpClientV1 {
	return NewEmailHttpClientV1WithConfig(nil)
}

func NewEmailHttpClientV1WithConfig(config *cconf.ConfigParams) *EmailHttpClientV1 {
	c := &EmailHttpClientV1{
		CommandableHttpClient: clients.NewCommandableHttpClient("v1/email"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *EmailHttpClientV1) SendMessage(ctx context.Context, correlationId string, message *EmailMessageV1, parameters *cconf.ConfigParams) error {
	params := cdata.NewAnyValueMapFromTuples(
		"message", message,
		"parameters", parameters,
	)

	_, err := c.CallCommand(ctx, "send_message", correlationId, params)

	return err
}

func (c *EmailHttpClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipient *EmailRecipientV1, message *EmailMessageV1, parameters *cconf.ConfigParams) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipient", recipient,
		"message", message,
		"parameters", parameters,
	)

	_, err := c.CallCommand(ctx, "send_message_to_recipient", correlationId, params)

	return err
}

func (c *EmailHttpClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*EmailRecipientV1, message *EmailMessageV1, parameters *cconf.ConfigParams) error {
	params := cdata.NewAnyValueMapFromTuples(
		"recipients", recipients,
		"message", message,
		"parameters", parameters,
	)

	_, err := c.CallCommand(ctx, "send_message_to_recipients", correlationId, params)

	return err
}
