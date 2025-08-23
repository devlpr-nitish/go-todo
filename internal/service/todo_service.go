package service

import (
	"github.com/devlpr-nitish/todo/internal/models"
	"github.com/devlpr-nitish/todo/internal/repository"
)


type TodoService struct{
	Repo *repository.TodoRepo
}

func NewTodoService(repo *repository.TodoRepo) *TodoService{
	return &TodoService{
		Repo: repo,
	}
}


func (s *TodoService) GetTodos() ([]models.Todo, error){
	return s.Repo.FindAll()
}