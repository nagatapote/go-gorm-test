package util

import (
	"regexp"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

type CustomValidator struct {
	Validator *validator.Validate
}

// PasswordValidate return bool
func PasswordValidate(fl validator.FieldLevel) bool {
	match1 := regexp.MustCompile(`[0-9]`).Match([]byte(fl.Field().String()))
	match2 := regexp.MustCompile(`[a-z]`).Match([]byte(fl.Field().String()))
	match3 := regexp.MustCompile(`[A-Z]`).Match([]byte(fl.Field().String()))
	return match1 == true && match2 == true && match3 == true
}

// Validate return err
func (cv *CustomValidator) Validate(i interface{}) error {
	cv.Validator.RegisterValidation("password", PasswordValidate)
	return cv.Validator.Struct(i)
}

// BindValidate return error
func BindValidate(c echo.Context, i interface{}) error {
	if err := c.Bind(i); err != nil {
		return err
	}
	if err := c.Validate(i); err != nil {
		return err
	}
	return nil
}
