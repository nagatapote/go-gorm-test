package route

import (
	"go-sqlboiler-test/interface/controllers"

	"github.com/labstack/echo"
)

type UserRouter interface {
	UserRouting(e *echo.Echo)
}

type userRouter struct {
	Uc controllers.UserController
}

func NewUserRouter(uc controllers.UserController) UserRouter {
	return userRouter{uc}
}

func (ur userRouter) UserRouting(e *echo.Echo) {
	r := e.Group("/user")
	r.GET("/", ur.Uc.UserGet)
	// r.POST("/")
	// r.PUT("/")
	// r.DELETE("/")
}
