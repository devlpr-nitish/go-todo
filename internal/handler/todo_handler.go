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

func (h *TodoHandler) Welcome(c *gin.Context){

	c.JSON(http.StatusOK, gin.H{"message":"Hello welcome to the API"});
}

func (h *TodoHandler) GetTodos(c *gin.Context){
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"user not found"})
		return
	}

	todos, err := h.Service.GetTodos(userID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not fetch todos"})
	}

	c.JSON(http.StatusOK, todos)
}

func (h *TodoHandler) CreateTodo(c *gin.Context){
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"user not found"})
		return
	}

	var todo models.Todo

	err := c.ShouldBindJSON(&todo)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Please send the valid data"})
		return
	}

	todo.UserID = userID.(uint)

	if err := h.Service.CreateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not create todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
} 


func (h *TodoHandler) GetTodoById(c *gin.Context){

	id, _ := strconv.Atoi(c.Param("id"))

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"user not found"})
		return
	}

	todo, err := h.Service.GetTodoById(uint(id), userID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"todo not found by id"})
		return
	}

	c.JSON(http.StatusOK, todo)
}




func (h *TodoHandler) UpdateTodo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))
	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"user not found"})
		return
	}

	todo , err := h.Service.GetTodoById(uint(id), userID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"todo not found by this id"})
		return
	}

	var req models.Todo

	if err := c.ShouldBindJSON(&req); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}

	todo.Title = req.Title
	todo.Completed = req.Completed

	err = h.Service.UpdateTodo(todo)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}


func (h *TodoHandler) DeleteTodo(c *gin.Context){
	id, _ := strconv.Atoi(c.Param("id"))

	userID, exists := c.Get("user_id")

	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error":"user not found"})
		return
	}

	todo , err := h.Service.GetTodoById(uint(id), userID.(uint))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"todo not found by this id"})
		return
	}

	if err := h.Service.DeleteTodo(todo.ID, userID.(uint)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not delete todo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"todo deleted"})
}