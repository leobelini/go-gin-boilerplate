package router

import (
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/router/user"

	"github.com/gin-gonic/gin"
)

func NewRouter(gin *gin.Engine, controllers *controller.Controller) {
	group := gin.Group("/api/v1")
	user.NewUserHandler(group, controllers)
}
