package repository

import "github.com/yusk/todo-sample-ddd-clean-architecture/domain/vo"

type SessionRepository interface {
	Get() (*vo.Session, error)
	Set(userID uint) error
	Clear() error
}
