package routers

import (
	user "leobelini/cashly/internal/routers/user"

	"github.com/gin-gonic/gin"
)

// @title           Meu SaaS API
// @version         1.0
// @description     API para meu sistema SaaS.
// @termsOfService  http://exemplo.com/terms/

// @contact.name   Suporte Técnico
// @contact.email  suporte@meusaas.com

// @host      localhost:8080
// @BasePath  /v1
func LoadRouters(r *gin.Engine) {
	routes := r.Group("/v1")

	user.LoadRouters(routes)
}
