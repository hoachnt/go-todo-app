package service

import (
	"github.com/hoachnt/go-todo-app/internal/domain"
	"github.com/hoachnt/go-todo-app/internal/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type TodoList interface {
	Create(userId int, list domain.TodoList) (int, error)
	GetAll(userId int) ([]domain.TodoList, error)
	GetById(userId, listId int) (domain.TodoList, error)
	Delete(userId, listId int) error
	Update(userId, listId int, input domain.UpdateListInput) error
}

type TodoItem interface {
	Create(userId, listId int, item domain.TodoItem) (int, error)
	GetAll(userId, listId int) ([]domain.TodoItem, error)
	GetById(userId, itemId int) (domain.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input domain.UpdateItemInput) error
}

type Service struct {
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
