package router

import (
	"github.com/devlpr-nitish/todo/internal/handler"
	"github.com/gin-gonic/gin"
)



func Setup(todoHandler *handler.TodoHandler) *gin.Engine{
	r := gin.Default();

	r.GET("/todos", todoHandler.GetTodos)


	return r
}