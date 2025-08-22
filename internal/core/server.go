package core

import (
	"fmt"
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/core/app"
	"leobelini/cashly/internal/core/dto"
	"leobelini/cashly/internal/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hibiken/asynq"
)

type AppServer struct {
	GinEngine *gin.Engine
	Database  *app.Database
	Env       *dto.DtoEnvApp
	Job       *asynq.Client
}

func NewAppServer() *AppServer {
	// Criar engine manualmente (gin.New) para controlar middleware
	engine := gin.New()

	baseApp := NewBaseApp()

	server := &AppServer{
		GinEngine: engine,
		Env:       baseApp.Env,
		Database:  baseApp.Database,
		Job:       baseApp.Job,
	}

	server.setupMiddleware()

	return server
}

// Configura todos os middlewares
func (s *AppServer) setupMiddleware() {

	if !s.Env.IsProd {
		startSwagger(s.GinEngine, s.Env)
		// CORS sempre aplicado antes das rotas
		s.GinEngine.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:3000"},
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
			AllowCredentials: true,
		}))
	}

	// Logger e Recovery sempre
	s.GinEngine.Use(gin.Logger())
	s.GinEngine.Use(gin.Recovery())

}

// Registra todas as rotas
func (s *AppServer) SetupRoutes(controllers *controller.Controller) {
	router.NewRouter(s.GinEngine, controllers) // passe controllers aqui
}

// Inicia o servidor
func (s *AppServer) StartServer() {
	url := fmt.Sprintf("%s:%d", s.Env.Server.Host, s.Env.Server.Port)
	fmt.Println("Server running at", url)

	if err := s.GinEngine.Run(url); err != nil {
		fmt.Println("Failed to start server:", err)
	}
}
