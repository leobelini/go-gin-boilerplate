package core

import (
	"fmt"
	_ "leobelini/cashly/docs"
	"leobelini/cashly/internal/core/dto"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func startSwagger(ginEngine *gin.Engine, env *dto.DtoEnvApp) {
	ginEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	url := fmt.Sprintf("%s:%d/swagger/index.html", env.Server.Host, env.Server.Port)

	fmt.Println("Swagger running at", url)
}
