package ports

import (
	"github.com/luschnat-ziegler/toDoListAPI/core/domain"
	"github.com/luschnat-ziegler/toDoListAPI/errs"
)

type ToDoListService interface {
	GetAllLists() (*[]domain.ToDoList, *errs.AppError)
	GetOneListById(string) (*domain.ToDoList, *errs.AppError)
	UpdateOneListById(string, domain.ToDoList) (*domain.ToDoList, *errs.AppError)
	SaveList(domain.ToDoList) (*domain.ToDoList, *errs.AppError)
	DeleteList(string) (*int64, *errs.AppError)
}