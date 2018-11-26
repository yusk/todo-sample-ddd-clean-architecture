package presenter

import (
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/output"
	"github.com/yusk/todo-sample-ddd-clean-architecture/usecase/input"
)

type SessionPresenter interface {
	PostSignUp(input.SessionInputPort) output.RedirectOutputPort
	PostSignIn(input.SessionInputPort) output.RedirectOutputPort
}

type SessionPresenterImpl struct {
}

func NewSessionPresenter() SessionPresenter {
	return SessionPresenterImpl{}
}

func (p SessionPresenterImpl) PostSignUp(sessionInput input.SessionInputPort) output.RedirectOutputPort {
	return output.RedirectOutputPort{
		Error: sessionInput.Error,
	}
}

func (p SessionPresenterImpl) PostSignIn(sessionInput input.SessionInputPort) output.RedirectOutputPort {
	return output.RedirectOutputPort{
		Error: sessionInput.Error,
	}
}
