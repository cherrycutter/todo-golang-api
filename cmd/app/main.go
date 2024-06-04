package main

import (
	_ "github.com/cherrycutter/todo_app/docs"
	"github.com/cherrycutter/todo_app/internal/app"
	"github.com/cherrycutter/todo_app/pkg/config"
	"github.com/cherrycutter/todo_app/pkg/logger"
)

// @Title Todo App API
// @version 1.0
// @description API Server for Todo App

// @host localhost:8080
// @BasePath /
func main() {
	logger.InitLogger()
	cfg := config.LoadConfig()
	app.Run(cfg)
}
