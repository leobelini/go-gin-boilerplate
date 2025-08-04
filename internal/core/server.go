package core

import (
	"fmt"
	"leobelini/cashly/internal/core/app"
	"leobelini/cashly/internal/core/dto"

	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

var AppEnv dto.DtoEnvApp

type AppServer struct {
	ginEngine *gin.Engine
	Database  *app.Database
	Env       *dto.DtoEnvApp
	Job       *asynq.Client
}

func NewAppServer(ginEngine *gin.Engine) *AppServer {
	baseApp := NewBaseApp()
	env := baseApp.Env
	dataBase := baseApp.Database

	job := app.NewJob(env)
	return &AppServer{ginEngine: ginEngine, Env: env, Database: dataBase, Job: job.Client}
}

func (s *AppServer) StartServer() {
	app.LoadEnv()

	if !s.Env.IsProd {
		startSwagger(s.ginEngine, s.Env)
		s.ginEngine.Use(gin.Logger())
	}

	s.ginEngine.Use(gin.Recovery())

	url := fmt.Sprintf("%s:%d", s.Env.Server.Host, s.Env.Server.Port)
	fmt.Println("Server running at", url)
	if err := s.ginEngine.Run(url); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
