package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	cconf "github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-commons-gox/data"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type EmailCommandableGrpcClientV1 struct {
	*clients.CommandableGrpcClient
}

func NewEmailCommandableGrpcClientV1() *EmailCommandableGrpcClientV1 {
	return NewEmailCommandableGrpcClientV1WithConfig(nil)
}

func NewEmailCommandableGrpcClientV1WithConfig(config *cconf.ConfigParams) *EmailCommandableGrpcClientV1 {
	c := &EmailCommandableGrpcClientV1{
		CommandableGrpcClient: clients.NewCommandableGrpcClient("v1/email"),
	}

	if config != nil {
		c.Configure(context.Background(), config)
	}

	return c
}

func (c *EmailCommandableGrpcClientV1) SendMessage(ctx context.Context, correlationId string, message *EmailMessageV1, parameters *config.ConfigParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "email.send_message")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_message", correlationId, data.NewAnyValueMapFromTuples(
		"message", message,
		"parameters", parameters,
	))

	return err
}

func (c *EmailCommandableGrpcClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipient *EmailRecipientV1, message *EmailMessageV1, parameters *config.ConfigParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "email.send_message_to_recipient")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_message_to_recipient", correlationId, data.NewAnyValueMapFromTuples(
		"recipient", recipient,
		"message", message,
		"parameters", parameters,
	))

	return err
}

func (c *EmailCommandableGrpcClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*EmailRecipientV1, message *EmailMessageV1, parameters *config.ConfigParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "email.send_message_to_recipients")
	defer timing.EndTiming(ctx, err)

	_, err = c.CallCommand(ctx, "send_message_to_recipients", correlationId, data.NewAnyValueMapFromTuples(
		"recipients", recipients,
		"message", message,
		"parameters", parameters,
	))

	return err
}
