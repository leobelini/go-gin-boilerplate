package dto

const TypeSendRecoveryPasswordEmail = "email:send_recovery_password"

type SendRecoveryPasswordPayload struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}
