package repository

import (
	"github.com/devlpr-nitish/todo/internal/models"
	"gorm.io/gorm"
)




type UserRepo struct{
	DB *gorm.DB
}


func NewUserRepo(db *gorm.DB) *UserRepo{
	return &UserRepo{DB:db}
}

func (r *UserRepo) Create(user models.User)error{
	return r.DB.Create(&user).Error
}

