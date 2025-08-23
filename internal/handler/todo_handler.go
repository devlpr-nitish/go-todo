package handler

import (
	"net/http"

	"github.com/devlpr-nitish/todo/internal/service"
	"github.com/gin-gonic/gin"
)



type TodoHandler struct{
	Service *service.TodoService
}


func NewTodoHandler(s *service.TodoService) *TodoHandler{
	return &TodoHandler{
		Service: s,
	}
}

func (h *TodoHandler) GetTodos(c *gin.Context){
	todos, err := h.Service.GetTodos()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not fetch todos"})
	}

	c.JSON(http.StatusOK, todos)
}