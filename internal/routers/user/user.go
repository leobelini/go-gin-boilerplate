package user

import "github.com/gin-gonic/gin"

func LoadRouters(routes *gin.RouterGroup) {
	group := routes.Group("/user")
	group.POST("/", createUser)
}
