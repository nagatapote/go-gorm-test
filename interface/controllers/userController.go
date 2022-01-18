package controllers

import (
	"go-gorm-test/domain/models"
	"go-gorm-test/usecase"
	"go-gorm-test/util"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type UserController interface {
	UserGet(c echo.Context) (err error)
	UserPost(c echo.Context) (err error)
	UserUpdate(c echo.Context) (err error)
	UserDelete(c echo.Context) (err error)
}

type userController struct {
	Cuu usecase.UserUseCase
}

func NewUserController(cuu usecase.UserUseCase) UserController {
	return userController{cuu}
}

type (
	userpost struct {
		Email          string `json:"email" validate:"required,email"`
		PasswordDigest string `json:"password" validate:"required,gte=8,password"`
	}

	userupdate struct {
		Email          string `json:"email" validate:"required,email"`
		PasswordDigest string `json:"password" validate:"required,gte=8,password"`
	}
)

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

func (uc userController) UserPost(c echo.Context) (err error) {
	// up := &userpost{
	// 	Email:          c.FormValue("email"),
	// 	PasswordDigest: c.FormValue("password"),
	// }

	//NOTE: body content type application/jsonであれば、post dataをこれで受け取れる。
	up := new(userpost)
	err = util.BindValidate(c, up)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	post, statuscode, err := uc.Cuu.UserPostUseCase(up.Email, up.PasswordDigest)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserUpdate(c echo.Context) (err error) {
	id := c.Param("id")
	uu := new(userupdate)
	err = util.BindValidate(c, uu)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	f, err := strconv.Atoi(id)
	post, statuscode, err := uc.Cuu.UserUpdateUseCase(f, uu.Email, uu.PasswordDigest)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (uc userController) UserDelete(c echo.Context) (err error) {
	id := c.Param("id")
	f, err := strconv.Atoi(id)
	post, statuscode, err := uc.Cuu.UserDeleteUseCase(f)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}
