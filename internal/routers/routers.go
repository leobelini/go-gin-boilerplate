package routers

import (
	user "leobelini/cashly/internal/routers/user"

	"github.com/gin-gonic/gin"
)

func LoadRouters(routes *gin.RouterGroup) {

	user.LoadRouters(routes)
}
