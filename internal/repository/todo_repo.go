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

func (r *TodoRepo) Create(todo models.Todo) error{
	return r.DB.Create(&todo).Error
}

func (r *TodoRepo) FindById(id uint) (models.Todo, error){
	var todo models.Todo
	result := r.DB.First(&todo, id)
	return todo, result.Error
}

func (r *TodoRepo) Update(todo models.Todo) error{
	return r.DB.Save(&todo).Error
}