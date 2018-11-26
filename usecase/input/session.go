package input

import "github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"

type SessionInputPort struct {
	User  *entity.User
	Error error
}
