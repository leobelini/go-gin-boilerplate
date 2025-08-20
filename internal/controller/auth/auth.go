package auth

import (
	internalDto "leobelini/cashly/internal/dto"
)

type AuthController struct {
	app *internalDto.DtoApp
}

func NewAuthController(app *internalDto.DtoApp) *AuthController {
	return &AuthController{app: app}
}
