package repository

import (
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"
)

type UserRepository interface {
	Get(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	AttachTodos(user *entity.User) (*entity.User, error)
	GetWithTodos(id uint) (*entity.User, error)
	Create(email string, password string) (*entity.User, error)
	AddToDo(user *entity.User, title string, content string) (*entity.User, *entity.Todo, error)
}
