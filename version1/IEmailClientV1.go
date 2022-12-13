package version1

import (
	"context"

	"github.com/pip-services3-gox/pip-services3-commons-gox/config"
)

type IEmailClientV1 interface {
	SendMessage(ctx context.Context, correlationId string, message *EmailMessageV1,
		parameters *config.ConfigParams) error

	SendMessageToRecipient(ctx context.Context, correlationId string, recipient *EmailRecipientV1,
		message *EmailMessageV1, parameters *config.ConfigParams) error

	SendMessageToRecipients(ctx context.Context, correlationId string, recipients []*EmailRecipientV1,
		message *EmailMessageV1, parameters *config.ConfigParams) error
}
