package user

import (
	userUseCase "leobelini/cashly/internal/usecase/user"

	"github.com/gin-gonic/gin"
)

func NewRoutersHandler(routerGroup *gin.RouterGroup, userUseCase *userUseCase.UserUseCase) {
	userRouter := routerGroup.Group("/user")

	userHandler := NewUserHandler(userUseCase)

	userRouter.POST("/", userHandler.CreateUser)
}
