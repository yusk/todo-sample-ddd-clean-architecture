package router

import (
	"html/template"
	"io"

	session "github.com/daisuke310vvv/echo-session"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/yusk/todo-sample-ddd-clean-architecture/adapter/handler"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Router(db *gorm.DB, store session.RedisStore) {
	tpl := &Template{
		templates: template.Must(template.ParseGlob("views/**/*.html")),
	}

	e := echo.New()
	e.Renderer = tpl

	e.Pre(middleware.MethodOverrideWithConfig(middleware.MethodOverrideConfig{
		Getter: middleware.MethodFromForm("_method"),
	}))
	e.Use(session.Sessions("session", store))
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLookup: "form:_csrf",
	}))

	sessionHandler := handler.NewSessionHandler(db)

	// e.GET("/sample/string", handler.SampleString)
	// e.GET("/sample/json", handler.SampleJSON)
	// e.GET("/sample/html", handler.SampleHTML)

	e.GET("/signup", sessionHandler.GetSignUp)
	e.POST("/signup", sessionHandler.PostSignUp)
	e.GET("/signin", sessionHandler.GetSignIn)
	e.GET("/signout", sessionHandler.GetSignOut)
	e.POST("/signin", sessionHandler.PostSignIn)

	// e.GET("/", handler.TodoList)
	// e.GET("/:id", handler.TodoShow)
	// e.POST("/", handler.TodoCreate)
	// e.GET("/new", handler.TodoNew)

	e.Start(":9090")
}
