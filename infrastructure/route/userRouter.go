package route

import (
	"go-gorm-test/interface/controllers"

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
	r.GET("/all", ur.Uc.UserGet)
	r.POST("/create", ur.Uc.UserPost)
	r.PUT("/update/:id", ur.Uc.UserUpdate)
	// r.DELETE("/delete", ur.Uc.UserDelete)
}
