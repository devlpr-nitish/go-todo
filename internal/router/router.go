package router

import (
	"github.com/devlpr-nitish/todo/internal/handler"
	"github.com/gin-gonic/gin"
)



func Setup(todoHandler *handler.TodoHandler, userHandler *handler.UserHandler) *gin.Engine{
	r := gin.Default();

	// todo routes
	r.GET("/", todoHandler.Welcome)
	r.GET("/todos", todoHandler.GetTodos)
	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos/:id", todoHandler.GetTodoById)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)

	// user routes
	r.POST("/users/register", userHandler.Register)


	return r
}