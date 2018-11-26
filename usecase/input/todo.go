package input

import "github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"

type TodoInputPort struct {
	User  *entity.User
	Error error
}

type TodoShowInputPort struct {
	User  *entity.User
	Todo  *entity.Todo
	Error error
}
