package version1

import (
	"context"

	"github.com/pip-services-infrastructure2/client-email-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
	"github.com/pip-services3-gox/pip-services3-grpc-gox/clients"
)

type EmailGrpcClientV1 struct {
	*clients.GrpcClient
}

func NewEmailGrpcClientV1() *EmailGrpcClientV1 {
	return &EmailGrpcClientV1{
		GrpcClient: clients.NewGrpcClient("email_v1.Email"),
	}
}

func (c *EmailGrpcClientV1) SendMessage(ctx context.Context, correlationId string,
	message *EmailMessageV1, parameters *config.ConfigParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "email_v1.send_message")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailMessageRequest{
		CorrelationId: correlationId,
		Message:       fromMessage(message),
	}

	if parameters != nil {
		req.Parameters = parameters.Value()
	}

	reply := new(protos.EmailEmptyReply)
	err = c.CallWithContext(ctx, "send_message", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *EmailGrpcClientV1) SendMessageToRecipient(ctx context.Context, correlationId string,
	recipient *EmailRecipientV1, message *EmailMessageV1,
	parameters *config.ConfigParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "email_v1.send_message_to_recipient")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailMessageWithRecipientRequest{
		CorrelationId: correlationId,
		Message:       fromMessage(message),
		Recipient:     fromRecipient(recipient),
	}

	if parameters != nil {
		req.Parameters = parameters.Value()
	}

	reply := new(protos.EmailEmptyReply)
	err = c.CallWithContext(ctx, "send_message_to_recipient", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}

func (c *EmailGrpcClientV1) SendMessageToRecipients(ctx context.Context, correlationId string,
	recipients []*EmailRecipientV1, message *EmailMessageV1,
	parameters *config.ConfigParams) (err error) {
	timing := c.Instrument(ctx, correlationId, "email_v1.send_message_to_recipients")
	defer timing.EndTiming(ctx, err)

	req := &protos.EmailMessageWithRecipientsRequest{
		CorrelationId: correlationId,
		Message:       fromMessage(message),
		Recipients:    fromRecipients(recipients),
	}

	if parameters != nil {
		req.Parameters = parameters.Value()
	}

	reply := new(protos.EmailEmptyReply)
	err = c.CallWithContext(ctx, "send_message_to_recipients", correlationId, req, reply)
	if err != nil {
		return err
	}

	if reply.Error != nil {
		err = toError(reply.Error)
		return err
	}

	return nil
}
