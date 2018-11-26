package handler

import (
	"fmt"
	"net/http"

	"github.com/daisuke310vvv/echo-session"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/database"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/handler/dto"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/presenter"
	"github.com/yusk/todo-sample-ddd-clean-architecture/domain/repository"
	"github.com/yusk/todo-sample-ddd-clean-architecture/usecase"
	validator "gopkg.in/go-playground/validator.v9"
)

type SessionHandler interface {
	GetSignUp(c echo.Context) error
	GetSignIn(c echo.Context) error
	GetSignOut(c echo.Context) error
	PostSignUp(c echo.Context) error
	PostSignIn(c echo.Context) error
}

type SessionHandlerImpl struct {
	validate         *validator.Validate
	userRepository   repository.UserRepository
	sessionPresenter presenter.SessionPresenter
}

func NewSessionHandler(db *gorm.DB) SessionHandler {
	return SessionHandlerImpl{
		validate:         validator.New(),
		userRepository:   database.NewUserRepository(db),
		sessionPresenter: presenter.NewSessionPresenter(),
	}
}

func (h SessionHandlerImpl) sessionRepository(c echo.Context) repository.SessionRepository {
	return database.NewSessionRepository(session.Default(c))
}

func (h SessionHandlerImpl) sessionUseCase(c echo.Context) usecase.SessionUseCase {
	return usecase.NewSessionUseCase(h.userRepository, h.sessionRepository(c))
}

func (h SessionHandlerImpl) GetSignUp(c echo.Context) error {
	mapData := map[string]interface{}{}
	mapData["CSRF"] = c.Get("csrf").(string)
	return c.Render(http.StatusFound, "session/signup", mapData)
}

func (h SessionHandlerImpl) GetSignOut(c echo.Context) error {
	sessionRepository := h.sessionRepository(c)
	err := sessionRepository.Clear()
	if err != nil {
		return c.Redirect(http.StatusFound, "/")
	}
	return c.Redirect(http.StatusFound, "/")
}

func (h SessionHandlerImpl) GetSignIn(c echo.Context) error {
	mapData := map[string]interface{}{}
	mapData["CSRF"] = c.Get("csrf").(string)
	return c.Render(http.StatusOK, "session/signin", mapData)
}

func (h SessionHandlerImpl) PostSignUp(c echo.Context) error {
	p, err := dto.ValidatedSessionParam(c, h.validate)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, c.Request().URL.String())
	}

	sessionUseCase := h.sessionUseCase(c)
	output := h.sessionPresenter.PostSignUp(sessionUseCase.SignUp(p.Email, p.Password))

	if output.Error != nil {
		fmt.Println(output.Error)
		return c.Redirect(http.StatusFound, "/signup")
	}
	return c.Redirect(http.StatusFound, "/")
}

func (h SessionHandlerImpl) PostSignIn(c echo.Context) error {
	p, err := dto.ValidatedSessionParam(c, h.validate)
	if err != nil {
		fmt.Println(err)
		return c.Redirect(http.StatusFound, c.Request().URL.String())
	}

	sessionUseCase := h.sessionUseCase(c)
	output := h.sessionPresenter.PostSignIn(sessionUseCase.SignIn(p.Email, p.Password))

	if output.Error != nil {
		fmt.Println(output.Error)
		return c.Redirect(http.StatusFound, "/signup")
	}
	return c.Redirect(http.StatusFound, "/")
}
