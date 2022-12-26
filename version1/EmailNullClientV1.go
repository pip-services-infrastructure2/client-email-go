package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type EmailNullClientV1 struct {
}

func NewEmailNullClientV1() *EmailNullClientV1 {
	return &EmailNullClientV1{}
}

func (c *EmailNullClientV1) SendMessage(ctx context.Context, correlationId string, message *EmailMessageV1, parameters *config.ConfigParams) error {
	return nil
}

func (c *EmailNullClientV1) SendMessageToRecipient(ctx context.Context, correlationId string, recipient *EmailRecipientV1, message *EmailMessageV1, parameters *config.ConfigParams) error {
	return nil
}

func (c *EmailNullClientV1) SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*EmailRecipientV1, message *EmailMessageV1, parameters *config.ConfigParams) error {
	return nil
}
