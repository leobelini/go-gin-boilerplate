package main

import (
	"fmt"
	"leobelini/cashly/config"
	"leobelini/cashly/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadServerEnv()
	env := config.GetServerEnv()

	r := gin.Default()
	routes := r.Group("/v1")

	routers.LoadRouters(routes)

	url := fmt.Sprintf("%s:%d", env.Host, env.Port)
	fmt.Println("Server running at", url)

	if err := r.Run(url); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
