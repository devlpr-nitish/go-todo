package main

import (
	"log"

	"github.com/devlpr-nitish/todo/internal/config"
	"github.com/devlpr-nitish/todo/internal/db"
	"github.com/devlpr-nitish/todo/internal/handler"
	"github.com/devlpr-nitish/todo/internal/repository"
	"github.com/devlpr-nitish/todo/internal/router"
	"github.com/devlpr-nitish/todo/internal/service"
)


func main(){

	cfg := config.Load();

	database := db.Connect(cfg)

	todoRepo := repository.NewTodoRepo(database)
	todoServices := service.NewTodoService(todoRepo)
	todoHandler := handler.NewTodoHandler(todoServices)

	r := router.Setup(todoHandler)

	log.Printf("App running on :%s", cfg.Port)

	if err := r.Run(":" + cfg.Port); err != nil{
		log.Fatal(err)
	}
}