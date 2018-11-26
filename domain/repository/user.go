package repository

import (
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"
)

type UserRepository interface {
	Get(id uint) (*entity.User, error)
	GetByEmail(email string) (*entity.User, error)
	Create(email string, password string) (*entity.User, error)
}
