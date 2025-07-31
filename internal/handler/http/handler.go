package http

import (
	"leobelini/cashly/internal/handler/http/user"
	"leobelini/cashly/internal/usecase"

	"github.com/gin-gonic/gin"
)

type RoutersHandler struct {
	// routerGroup *gin.RouterGroup
	useCase *usecase.UseCase
}

// @title           Meu SaaS API
// @version         1.0
// @description     API para meu sistema SaaS.
// @termsOfService  http://exemplo.com/terms/

// @contact.name   Suporte Técnico
// @contact.email  suporte@meusaas.com

// @host      localhost:8080
// @BasePath  /v1

func NewRoutersHandler(useCase *usecase.UseCase) *RoutersHandler {
	return &RoutersHandler{useCase: useCase}
}

func (h *RoutersHandler) LoadRouters(r *gin.Engine) {
	routes := r.Group("/v1")

	user.NewRoutersHandler(routes, h.useCase.User)
}
