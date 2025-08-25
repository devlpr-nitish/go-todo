package repository

import (
	"fmt"
	"log"

	"github.com/devlpr-nitish/todo/internal/models"
	"github.com/devlpr-nitish/todo/internal/utils"
	"golang.org/x/crypto/bcrypt"
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

func (r *UserRepo) FindByEmail(email string) (models.User, error){
	var user models.User

	result := r.DB.Where("email = ?", email).First(&user)

	return user, result.Error
}

func CheckPassword(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

func (r *UserRepo) FindById(id uint) (models.User, error){
	var user models.User
	result := r.DB.First(&user, id)

	return user, result.Error
}


func (r *UserRepo) Login(req models.LoginReq) (string, error){
	user, err := r.FindByEmail(req.Email)

	if err != nil{
		return "", err
	}

	log.Printf("User found: %+v", user)

	if CheckPassword(req.Password, user.Password) {
		return "", fmt.Errorf("invalid email or password")
	}

	token, err := utils.GenerateJwtSecret(user.ID, user.Email)


	if err != nil{
		return "", err
	}

	return token, nil
}