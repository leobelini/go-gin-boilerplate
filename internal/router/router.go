package router

import (
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/router/user"

	"github.com/gin-gonic/gin"
)

// @title           Meu SaaS API
// @version         1.0
// @description     API para meu sistema SaaS.
// @termsOfService  http://exemplo.com/terms/

// @contact.name   Suporte Técnico
// @contact.email  suporte@meusaas.com

// @host      localhost:8080
// @BasePath  /api/v1
func NewRouter(gin *gin.Engine, controllers *controller.Controller) {
	group := gin.Group("/api/v1")
	user.NewUserHandler(group, controllers)
}
