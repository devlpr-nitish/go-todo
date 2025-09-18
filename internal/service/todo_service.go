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


func (s *TodoService) GetTodos(userID uint) ([]models.Todo, error){
	return s.Repo.FindByUserID(userID)
}

func (s *TodoService) CreateTodo(todo models.Todo) error{
	return s.Repo.Create(todo)
}


func (s *TodoService) GetTodoById(id uint, userID uint) (models.Todo, error){
	return s.Repo.FindByIdAndUser(id, userID)
}

func (s *TodoService) UpdateTodo(todo models.Todo) error{
	return s.Repo.Update(todo)
}

func (s *TodoService) DeleteTodo(id uint, userID uint) error{
	return s.Repo.Delete(id, userID)
}