package auth

import (
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/middleware"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	controllers *controller.Controller
}

func NewAuthHandler(group *gin.RouterGroup, middleware *middleware.MiddlewareHandler, controllers *controller.Controller) {
	router := group.Group("/auth")

	handler := &AuthHandler{controllers: controllers}

	router.PUT("/sign-up-confirm/:token", handler.SignUpConfirm)
	router.POST("/password-recovery", handler.PasswordRecovery)
	router.PUT("/reset-password/:token", handler.ResetPassword)
	router.POST("/sign-in", handler.SignIn)
	router.Use(middleware.JWTAuthMiddleware()).GET("/me", handler.Me)
}
