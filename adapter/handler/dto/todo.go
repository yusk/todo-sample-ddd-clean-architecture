package dto

import (
	"github.com/labstack/echo"
	validator "gopkg.in/go-playground/validator.v9"
)

type todoParam struct {
	Title   string `json:"email" validate:"required"`
	Content string `json:"password"`
}

func ValidatedTodoParam(c echo.Context, v *validator.Validate) (*todoParam, error) {
	var p todoParam
	err := c.Bind(&p)
	if err != nil {
		return nil, err
	}
	err = v.Struct(p)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
