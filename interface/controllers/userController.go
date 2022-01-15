package controllers

import (
	"go-sqlboiler-test/domain/models"
	"go-sqlboiler-test/usecase"
	"net/http"

	"github.com/labstack/echo"
)

type UserController interface {
	UserGet(c echo.Context) (err error)
	// UserPost(c echo.Context) (err error)
	// UserPut(c echo.Context) (err error)
	// UserDelete(c echo.Context) (err error)
}

type userController struct {
	Cuu usecase.UserUseCase
}

func NewUserController(cuu usecase.UserUseCase) UserController {
	return userController{cuu}
}

func (uc userController) UserGet(c echo.Context) (err error) {
	post, statuscode, err := uc.Cuu.UserGetUseCase()
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}
