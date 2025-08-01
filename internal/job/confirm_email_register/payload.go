package confirmemailregister

const TypeSendConfirmationEmail = "email:send_confirmation_email"

type Payload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
