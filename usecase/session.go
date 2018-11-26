package usecase

import (
	"github.com/pkg/errors"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/repository"
	"github.com/yusk/todo-sample-ddd-clean-architecture/usecase/input"
)

type SessionUseCase interface {
	SignUp(email string, password string) input.SessionInputPort
	SignIn(email string, password string) input.SessionInputPort
}

type SessionUseCaseImpl struct {
	UserRepository    repository.UserRepository
	SessionRepository repository.SessionRepository
}

func NewSessionUseCase(userRepository repository.UserRepository, sessionRepository repository.SessionRepository) SessionUseCase {
	return SessionUseCaseImpl{
		UserRepository:    userRepository,
		SessionRepository: sessionRepository,
	}
}

func (uc SessionUseCaseImpl) SignUp(email string, password string) input.SessionInputPort {
	user, err := uc.UserRepository.Create(email, password)
	if err != nil {
		return input.SessionInputPort{
			User:  nil,
			Error: err,
		}
	}

	err = uc.SessionRepository.Set(user.ID)
	return input.SessionInputPort{
		User:  user,
		Error: err,
	}
}

func (uc SessionUseCaseImpl) SignIn(email string, password string) input.SessionInputPort {
	user, err := uc.UserRepository.GetByEmail(email)
	if err != nil {
		return input.SessionInputPort{
			User:  nil,
			Error: err,
		}
	}

	if !user.IsAuthenticated(password) {
		return input.SessionInputPort{
			User:  nil,
			Error: errors.Errorf("No Authenticated"),
		}
	}

	err = uc.SessionRepository.Set(user.ID)
	return input.SessionInputPort{
		User:  user,
		Error: err,
	}
}
