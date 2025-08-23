package repository

import (
	"github.com/devlpr-nitish/todo/internal/models"
	"gorm.io/gorm"
)



type TodoRepo struct{
	DB *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo{
	return &TodoRepo{
		DB: db,
	}
}

func (r *TodoRepo) FindAll() ([]models.Todo, error){
	var todos []models.Todo
	result := r.DB.Find(&todos);
	return todos, result.Error
}
