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

func (r *TodoRepo) FindByUserID(userID uint) ([]models.Todo, error){
	var todos []models.Todo
	result := r.DB.Where("user_id = ? ", userID).Find(&todos);
	return todos, result.Error
}

func (r *TodoRepo) Create(todo models.Todo) error{
	return r.DB.Create(&todo).Error
}

func (r *TodoRepo) FindByIdAndUser(id uint, userID uint) (models.Todo, error){
	var todo models.Todo
	result := r.DB.Where("id = ? AND user_id = ?", id, userID).First(&todo)
	return todo, result.Error
}

func (r *TodoRepo) Update(todo models.Todo) error{
	return r.DB.Save(&todo).Error
}

func (r *TodoRepo) Delete(id uint, userID uint) error{
	return r.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Todo{}).Error
}