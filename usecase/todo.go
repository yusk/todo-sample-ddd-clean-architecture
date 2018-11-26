package usecase

import (
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/entity"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/repository"
	"github.com/yusk/todo-sample-ddd-clean-architecture/usecase/input"
)

type TodoUseCase interface {
	List() input.TodoInputPort
	Show(id uint) input.TodoShowInputPort
	New() input.TodoInputPort
	Create(title string, content string) input.TodoShowInputPort
}

type TodoUseCaseImpl struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
}

func NewTodoUseCase(userRepository repository.UserRepository, sessionRepository repository.SessionRepository) TodoUseCase {
	return TodoUseCaseImpl{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
}

func (r TodoUseCaseImpl) user() (*entity.User, error) {
	sess, err := r.SessionRepository.Get()
	if err != nil {
		return nil, err
	}

	user, err := r.UserRepository.Get(sess.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r TodoUseCaseImpl) userWithTodo() (*entity.User, error) {
	sess, err := r.SessionRepository.Get()
	if err != nil {
		return nil, err
	}

	user, err := r.UserRepository.GetWithTodos(sess.UserID)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r TodoUseCaseImpl) List() input.TodoInputPort {
	user, err := r.userWithTodo()
	if err != nil {
		return input.TodoInputPort{
			User:  nil,
			Error: err,
		}
	}
	return input.TodoInputPort{
		User:  user,
		Error: nil,
	}
}

func (r TodoUseCaseImpl) Show(id uint) input.TodoShowInputPort {
	user, err := r.userWithTodo()
	if err != nil {
		return input.TodoShowInputPort{
			User:  nil,
			Todo:  nil,
			Error: err,
		}
	}

	todo, err := user.GetTodo(id)
	if err != nil {
		return input.TodoShowInputPort{
			User:  nil,
			Todo:  nil,
			Error: err,
		}
	}

	return input.TodoShowInputPort{
		User:  user,
		Todo:  todo,
		Error: nil,
	}
}

func (r TodoUseCaseImpl) New() input.TodoInputPort {
	user, err := r.user()
	if err != nil {
		return input.TodoInputPort{
			User:  nil,
			Error: err,
		}
	}

	return input.TodoInputPort{
		User:  user,
		Error: nil,
	}
}

func (r TodoUseCaseImpl) Create(title string, content string) input.TodoShowInputPort {
	user, err := r.user()
	if err != nil {
		return input.TodoShowInputPort{
			User:  nil,
			Todo:  nil,
			Error: err,
		}
	}

	user, todo, err := r.UserRepository.AddToDo(user, title, content)
	if err != nil {
		return input.TodoShowInputPort{
			User:  nil,
			Todo:  nil,
			Error: err,
		}
	}

	return input.TodoShowInputPort{
		User:  user,
		Todo:  todo,
		Error: nil,
	}
}
