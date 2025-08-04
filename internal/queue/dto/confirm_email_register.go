package dto

const TypeSendConfirmationEmail = "email:send_confirmation_email"

type SendConfirmationEmailRegisterPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
