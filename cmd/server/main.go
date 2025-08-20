package main

import (
	"leobelini/cashly/internal/controller"
	"leobelini/cashly/internal/core"
	"leobelini/cashly/internal/entity"
	"leobelini/cashly/internal/model"
	"leobelini/cashly/internal/queue/job"
	"leobelini/cashly/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	server := core.NewAppServer(r)

	// Load database
	server.Database.Start()
	defer server.Database.Close()

	if server.Env.Database.AutoMigrate {
		server.Database.Migrate(&entity.User{})
	}

	job := job.NewJob(server.Job)

	models := model.LoadModels(server.Database.Db)
	controllers := controller.NewController(models, job, server.Env)

	router.NewRouter(r, controllers)

	server.StartServer()
}
