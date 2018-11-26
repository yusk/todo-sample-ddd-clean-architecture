package handler

import (
	"fmt"
	"net/http"
	"strconv"

	session "github.com/daisuke310vvv/echo-session"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/database"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/handler/dto"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/presenter"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/repository"
	"github.com/yusk/todo-sample-ddd-clean-architecture/usecase"
	validator "gopkg.in/go-playground/validator.v9"
)

type TodoHandler struct {
	validate       *validator.Validate
	userRepository repository.UserRepository
	todoPresenter  presenter.TodoPresenter
}

func NewTodoHandler(db *gorm.DB) TodoHandler {
	return TodoHandler{
		validate:       validator.New(),
		userRepository: database.NewUserRepository(db),
		todoPresenter:  presenter.NewTodoPresenter(),
	}
}

func (h TodoHandler) sessionRepository(c echo.Context) repository.SessionRepository {
	return database.NewSessionRepository(session.Default(c))
}

func (h TodoHandler) todoUseCase(c echo.Context) usecase.TodoUseCase {
	return usecase.NewTodoUseCase(h.userRepository, h.sessionRepository(c))
}

func (h TodoHandler) List(c echo.Context) error {
	todoUseCase := h.todoUseCase(c)
	output := h.todoPresenter.List(todoUseCase.List())

	if output.Error != nil {
		return c.Redirect(http.StatusFound, "/signin")
	}
	return c.Render(http.StatusFound, "todo/list", output.Context)
}

func (h TodoHandler) Show(c echo.Context) error {
	idStr := c.Param("id")
	failureURL := "/"

	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, failureURL)
	}
	id := uint(id64)

	todoUseCase := h.todoUseCase(c)
	output := h.todoPresenter.Show(todoUseCase.Show(id))

	if output.Error != nil {
		return c.Redirect(http.StatusFound, failureURL)
	}
	return c.Render(http.StatusOK, "todo/show", output.Context)
}

func (h TodoHandler) New(c echo.Context) error {
	todoUseCase := h.todoUseCase(c)
	output := h.todoPresenter.New(todoUseCase.New())

	if output.Error != nil {
		return c.Redirect(http.StatusFound, "/")
	}

	output.Context["CSRF"] = c.Get("csrf").(string)
	return c.Render(http.StatusOK, "todo/new", output.Context)
}

func (h TodoHandler) Create(c echo.Context) error {
	p, err := dto.ValidatedTodoParam(c, h.validate)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, c.Request().URL.String())
	}

	todoUseCase := h.todoUseCase(c)
	output := h.todoPresenter.Create(todoUseCase.Create(p.Title, p.Content))

	if output.Error != nil {
		return c.Redirect(http.StatusFound, "/new")
	}
	return c.Redirect(http.StatusFound, "/")
}
