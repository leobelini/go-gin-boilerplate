package main

import (
	"fmt"
	"leobelini/cashly/config"
	docs "leobelini/cashly/docs"
	"leobelini/cashly/internal/routers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadServerEnv()
	env := config.GetServerEnv()

	r := gin.Default()
	routers.LoadRouters(r)

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	url := fmt.Sprintf("%s:%d", env.Host, env.Port)
	fmt.Println("Server running at", url)

	if err := r.Run(url); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
