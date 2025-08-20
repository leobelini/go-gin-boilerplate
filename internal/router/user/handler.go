package user

import (
	"leobelini/cashly/internal/controller"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	controllers *controller.Controller
}

func NewUserHandler(group *gin.RouterGroup, controllers *controller.Controller) {
	router := group.Group("/user")

	handler := &UserHandler{controllers: controllers}

	router.POST("/", handler.CreateUser)
}
