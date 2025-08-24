package handler

import (
	"net/http"

	"github.com/devlpr-nitish/todo/internal/models"
	"github.com/devlpr-nitish/todo/internal/service"
	"github.com/gin-gonic/gin"
)




type UserHandler struct{
	Service *service.UserService
}


func NewUserHandler(s *service.UserService) *UserHandler{
	return &UserHandler{Service: s,}
}


func (h *UserHandler) Register(c *gin.Context){
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error":"invalid request"})
		return
	}

	if err := h.Service.Register(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"could not register user"})
		return
	}

	c.JSON(http.StatusOK, user)
}