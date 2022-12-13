package version1

import (
	"encoding/json"

	"github.com/pip-services-infrastructure2/client-email-go/protos"
	"github.com/pip-services3-gox/pip-services3-commons-gox/convert"
	"github.com/pip-services3-gox/pip-services3-commons-gox/errors"
)

func fromError(err error) *protos.ErrorDescription {
	if err == nil {
		return nil
	}

	desc := errors.ErrorDescriptionFactory.Create(err)
	obj := &protos.ErrorDescription{
		Type:          desc.Type,
		Category:      desc.Category,
		Code:          desc.Code,
		CorrelationId: desc.CorrelationId,
		Status:        convert.StringConverter.ToString(desc.Status),
		Message:       desc.Message,
		Cause:         desc.Cause,
		StackTrace:    desc.StackTrace,
		Details:       fromMap(desc.Details),
	}

	return obj
}

func toError(obj *protos.ErrorDescription) error {
	if obj == nil || (obj.Category == "" && obj.Message == "") {
		return nil
	}

	description := &errors.ErrorDescription{
		Type:          obj.Type,
		Category:      obj.Category,
		Code:          obj.Code,
		CorrelationId: obj.CorrelationId,
		Status:        convert.IntegerConverter.ToInteger(obj.Status),
		Message:       obj.Message,
		Cause:         obj.Cause,
		StackTrace:    obj.StackTrace,
		Details:       toMap(obj.Details),
	}

	return errors.ApplicationErrorFactory.Create(description)
}

func fromMap(val map[string]any) map[string]string {
	r := map[string]string{}

	for k, v := range val {
		r[k] = convert.StringConverter.ToString(v)
	}

	return r
}

func toMap(val map[string]string) map[string]any {
	r := map[string]any{}

	for k, v := range val {
		r[k] = v
	}

	return r
}

func toJson(value any) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func fromJson(value string) any {
	if value == "" {
		return nil
	}

	var m any
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromStringMap(value map[string]string) string {
	if value == nil {
		return ""
	}

	b, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(b[:])
}

func toStringMap(value string) map[string]string {
	if value == "" {
		return nil
	}

	var m map[string]string
	json.Unmarshal([]byte(value), &m)
	return m
}

func fromMessage(message *EmailMessageV1) *protos.EmailMessage {
	if message == nil {
		return nil
	}

	obj := &protos.EmailMessage{
		From:    message.From,
		Cc:      message.Cc,
		Bcc:     message.Bcc,
		To:      message.To,
		ReplyTo: message.ReplyTo,
		Subject: fromStringMap(message.Subject),
		Text:    fromStringMap(message.Text),
		Html:    fromStringMap(message.Html),
	}

	return obj
}

func toMessage(obj *protos.EmailMessage) *EmailMessageV1 {
	if obj == nil {
		return nil
	}

	message := &EmailMessageV1{
		From:    obj.From,
		Cc:      obj.Cc,
		Bcc:     obj.Bcc,
		To:      obj.To,
		ReplyTo: obj.ReplyTo,
		Subject: toStringMap(obj.Subject),
		Text:    toStringMap(obj.Text),
		Html:    toStringMap(obj.Html),
	}

	return message
}

func fromRecipient(recipient *EmailRecipientV1) *protos.EmailRecipient {
	if recipient == nil {
		return nil
	}

	obj := &protos.EmailRecipient{
		Id:       recipient.Id,
		Name:     recipient.Name,
		Email:    recipient.Email,
		Language: recipient.Language,
	}

	return obj
}

func toRecipient(obj *protos.EmailRecipient) *EmailRecipientV1 {
	if obj == nil {
		return nil
	}

	recipient := &EmailRecipientV1{
		Id:       obj.Id,
		Name:     obj.Name,
		Email:    obj.Email,
		Language: obj.Language,
	}

	return recipient
}

func fromRecipients(recipients []*EmailRecipientV1) []*protos.EmailRecipient {
	if recipients == nil {
		return nil
	}

	data := make([]*protos.EmailRecipient, len(recipients))

	for i, v := range recipients {
		data[i] = fromRecipient(v)
	}

	return data
}

func toRecipients(obj []*protos.EmailRecipient) []*EmailRecipientV1 {
	if obj == nil {
		return nil
	}

	data := make([]*EmailRecipientV1, len(obj))

	for i, v := range obj {
		data[i] = toRecipient(v)
	}

	return data
}
