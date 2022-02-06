package route

import (
	"go-gorm-test/interface/controllers"

	"github.com/labstack/echo"
)

type IndexRouter interface {
	Routing(e *echo.Echo)
}

type indexRouter struct {
	Ur UserRouter
	Fr FileRouter
}

func NewIndexRouter(ur UserRouter, fr FileRouter) IndexRouter {
	return indexRouter{ur, fr}
}

func (ir indexRouter) Routing(e *echo.Echo) {
	e.GET("/", controllers.Root)
	ir.Ur.UserRouting(e)
}
