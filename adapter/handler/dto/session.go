package dto

import (
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

type sessionParam struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

func ValidatedSessionParam(c echo.Context, v *validator.Validate) (*sessionParam, error) {
	var p sessionParam
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
