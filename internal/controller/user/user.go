package user

import (
	internalDto "leobelini/cashly/internal/dto"
)

type UserController struct {
	app *internalDto.DtoApp
}

func NewUserController(app *internalDto.DtoApp) *UserController {
	return &UserController{app: app}
}
