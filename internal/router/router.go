package router

import (
	"github.com/devlpr-nitish/todo/internal/handler"
	"github.com/gin-gonic/gin"
)



func Setup(todoHandler *handler.TodoHandler) *gin.Engine{
	r := gin.Default();

	r.GET("/", todoHandler.Welcome)
	r.GET("/todos", todoHandler.GetTodos)
	r.POST("/todos", todoHandler.CreateTodo)
	r.GET("/todos/:id", todoHandler.GetTodoById)
	r.PUT("/todos/:id", todoHandler.UpdateTodo)
	r.DELETE("/todos/:id", todoHandler.DeleteTodo)


	return r
}