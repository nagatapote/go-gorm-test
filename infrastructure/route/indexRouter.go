package route

import (
	"go-sqlboiler-test/interface/controllers"

	"github.com/labstack/echo"
)

type IndexRouter interface {
	Routing(e *echo.Echo)
}

type indexRouter struct {
	Ur UserRouter
}

func NewIndexRouter(ur UserRouter) IndexRouter {
	return indexRouter{ur}
}

func (ir indexRouter) Routing(e *echo.Echo) {
	e.GET("/", controllers.Root)
	ir.Ur.UserRouting(e)
}