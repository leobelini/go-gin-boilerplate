package core

import (
	"fmt"
	"leobelini/cashly/internal/core/app"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startSwagger(ginEngine *gin.Engine) {
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	env := app.AppEnv
	url := fmt.Sprintf("%s:%d/swagger/index.html", env.Server.Host, env.Server.Port)

	fmt.Println("Swagger running at", url)
}
