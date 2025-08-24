package middleware

import "leobelini/cashly/internal/controller"

type MiddlewareHandler struct {
	controllers *controller.Controller
}

func NewMiddlewareHandler(controllers *controller.Controller) *MiddlewareHandler {
	return &MiddlewareHandler{controllers: controllers}
}
