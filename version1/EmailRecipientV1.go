package version1

type EmailRecipientV1 struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Language string `json:"language"`
}

func EmptyEmailRecipientV1() *EmailRecipientV1 {
	return &EmailRecipientV1{}
}

func NewEmailRecipientV1(id string, name string, email string, language string) *EmailRecipientV1 {
	return &EmailRecipientV1{
		Id:       id,
		Name:     name,
		Email:    email,
		Language: language,
	}
}
