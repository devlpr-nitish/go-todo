package models

import "gorm.io/gorm"


type User struct{
	gorm.Model
	Name string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Password string `json:"-"`
}


type LoginReq struct{
	Email string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}