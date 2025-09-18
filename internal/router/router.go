package router

import (
	"github.com/devlpr-nitish/todo/internal/handler"
	"github.com/devlpr-nitish/todo/internal/middlewares"
	"github.com/gin-gonic/gin"
)

func Setup(todoHandler *handler.TodoHandler, userHandler *handler.UserHandler) *gin.Engine {
	r := gin.Default()

	// todo routes
	r.GET("/", todoHandler.Welcome)
	r.GET("/todos", middlewares.AuthMiddleware(), todoHandler.GetTodos)
	r.POST("/todos", middlewares.AuthMiddleware(), todoHandler.CreateTodo)
	r.GET("/todos/:id", middlewares.AuthMiddleware(), todoHandler.GetTodoById)
	r.PUT("/todos/:id", middlewares.AuthMiddleware(), todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", middlewares.AuthMiddleware(), todoHandler.DeleteTodo)

	// user routes
	r.POST("/users/register", userHandler.Register)
	r.POST("/users/login", userHandler.Login)

	return r
}
