package app

import (
	"context"
	"github.com/cherrycutter/todo_app/internal/db"
	"github.com/cherrycutter/todo_app/internal/handlers"
	"github.com/cherrycutter/todo_app/internal/repos"
	"github.com/cherrycutter/todo_app/internal/services"
	"github.com/cherrycutter/todo_app/pkg/config"
	"github.com/cherrycutter/todo_app/pkg/logger"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Run(cfg config.Config) {
	database, err := db.InitDB(cfg)
	if err != nil {
		logger.Error.Fatal(err)
	}
	logger.Info.Printf("database connection established successfully")

	defer database.Close(context.Background())

	r := gin.Default()

	repo := repos.NewTodoRepo(database)
	service := services.NewTodoService(repo)
	handler := handlers.NewTodoHandler(service)

	handler.RegisterRoutes(r)

	// swagger endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	logger.Info.Println("starting server...")
	if err = r.Run(":8080"); err != nil {
		logger.Error.Fatal(err)
	}
}
