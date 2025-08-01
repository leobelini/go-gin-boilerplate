package main

import (
	"fmt"
	"leobelini/cashly/config"
	docs "leobelini/cashly/docs"
	"leobelini/cashly/internal/handler/http"
	"leobelini/cashly/internal/infra"
	"leobelini/cashly/internal/integration"
	"leobelini/cashly/internal/usecase"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadServerEnv()
	env := config.GetServerEnv()

	// Load database
	dbGorm, err := integration.StartDatabase()
	if err != nil {
		panic(err)
	}

	// Load Job
	job := integration.StartJob()

	gormRepository := infra.NewRepositoryGorm(dbGorm)
	useCaseApp := usecase.NewUseCase(gormRepository)

	r := gin.Default()

	handler := http.NewRoutersHandler(useCaseApp)
	handler.LoadRouters(r)

	docs.SwaggerInfo.BasePath = "/v1"
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	url := fmt.Sprintf("%s:%d", env.Host, env.Port)
	fmt.Println("Server running at", url)

	if err := r.Run(url); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
