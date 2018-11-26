package presenter

import (
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/output"
	"github.com/yusk/todo-sample-ddd-clean-architecture/usecase/input"
)

type TodoPresenter interface {
	List(i input.TodoInputPort) output.RenderOutputPort
	Show(i input.TodoShowInputPort) output.RenderOutputPort
	New(i input.TodoInputPort) output.RenderOutputPort
	Create(i input.TodoShowInputPort) output.RedirectOutputPort
}

type TodoPresenterImpl struct {
}

func NewTodoPresenter() TodoPresenter {
	return TodoPresenterImpl{}
}

func (p TodoPresenterImpl) List(i input.TodoInputPort) output.RenderOutputPort {
	context := map[string]interface{}{}
	context["User"] = i.User
	return output.RenderOutputPort{
		Context: context,
		Error:   i.Error,
	}
}

func (p TodoPresenterImpl) Show(i input.TodoShowInputPort) output.RenderOutputPort {
	context := map[string]interface{}{}
	context["User"] = i.User
	context["Todo"] = i.Todo
	return output.RenderOutputPort{
		Context: context,
		Error:   i.Error,
	}
}

func (p TodoPresenterImpl) New(i input.TodoInputPort) output.RenderOutputPort {
	context := map[string]interface{}{}
	context["User"] = i.User
	return output.RenderOutputPort{
		Context: context,
		Error:   i.Error,
	}
}

func (p TodoPresenterImpl) Create(i input.TodoShowInputPort) output.RedirectOutputPort {
	return output.RedirectOutputPort{
		Error: i.Error,
	}
}
