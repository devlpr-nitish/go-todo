package handler

import (
	"net/http"
	"strconv"

	"github.com/devlpr-nitish/todo/internal/models"
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

func (h *TodoHandler) CreateTodo(c *gin.Context){
	var todo models.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}

	if err := h.Service.CreateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not create todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
} 


func (h *TodoHandler) GetTodoById(c *gin.Context){

	id, _ := strconv.Atoi(c.Param("id"))

	todo, err := h.Service.GetTodoById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"todo not found by this id"})
		return
	}

	c.JSON(http.StatusOK, todo)
}


func (h *TodoHandler) Welcome(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{"message":"Hello welcome to the API"});
}

func (h *TodoHandler) UpdateTodo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	todo , err := h.Service.GetTodoById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"todo not found by this id"})
		return
	}

	if err := c.ShouldBindJSON(&todo); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}

	err = h.Service.UpdateTodo(todo)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}


func (h *TodoHandler) DeleteTodo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	_ , err := h.Service.GetTodoById(uint(id))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"todo not found by this id"})
		return
	}

	if err := h.Service.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"todo deleted"})
}