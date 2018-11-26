package database

import (
	session "github.com/daisuke310vvv/echo-session"
	"github.com/pkg/errors"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/repository"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/vo"
)

type SessionRepositoryImpl struct {
	session session.Session
}

func NewSessionRepository(session session.Session) repository.SessionRepository {
	return SessionRepositoryImpl{
		session: session,
	}
}

func (r SessionRepositoryImpl) Get() (*vo.Session, error) {
	v := r.session.Get("session")
	if v == nil {
		return nil, errors.Errorf("Session Not Found")
	}
	sess, ok := v.(vo.Session)
	if !ok {
		return nil, errors.Errorf("Session Cast Error")
	}
	return &sess, nil
}

func (r SessionRepositoryImpl) Set(userID uint) error {
	sess := vo.Session{UserID: userID}
	r.session.Set("session", &sess)
	return r.session.Save()
}

func (r SessionRepositoryImpl) Clear() error {
	r.session.Clear()
	r.session.Save()
	return nil
}
