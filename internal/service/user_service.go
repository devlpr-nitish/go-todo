package service

import (
	"github.com/devlpr-nitish/todo/internal/models"
	"github.com/devlpr-nitish/todo/internal/repository"
	"golang.org/x/crypto/bcrypt"
)


type UserService struct{
	Repo *repository.UserRepo
}


func NewUserService(repo *repository.UserRepo) *UserService{
	return &UserService{Repo: repo}
}


func (s *UserService) Register(user models.User) error{

	// hash password
	hashedPassword , err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil{
		return err
	}

	user.Password = string(hashedPassword)
	return s.Repo.Create(user)
}

func (s *UserService) Login(req models.LoginReq) (string, error){
	return s.Repo.Login(req)
}