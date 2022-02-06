package route

import (
	"go-gorm-test/interface/controllers"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type FileRouter interface {
	FileRouting(e *echo.Echo)
	fileCertificationRouting(e *echo.Echo)
}

type fileRouter struct {
	Fc controllers.FileController
}

func NewFileRouter(fc controllers.FileController) FileRouter {
	return fileRouter{fc}
}

func (fr fileRouter) FileRouting(e *echo.Echo) {
	fr.fileCertificationRouting(e)
}

func (fr fileRouter) fileCertificationRouting(e *echo.Echo) {
	r := e.Group("/files")
	r.Use(middleware.JWT([]byte(os.Getenv("SIGNINGKEY"))))
	r.POST("/upload", fr.Fc.FileUpload)
	r.POST("/download", fr.Fc.FileDownload)
}
