package auth

import (
	"leobelini/cashly/internal/controller"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	controllers *controller.Controller
}

func NewAuthHandler(group *gin.RouterGroup, controllers *controller.Controller) {
	router := group.Group("/auth")

	handler := &AuthHandler{controllers: controllers}

	router.PUT("/sign-up-confirm/:token", handler.SignUpConfirm)
}
