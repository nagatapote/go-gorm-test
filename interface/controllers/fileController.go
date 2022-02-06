package controllers

import (
	"go-gorm-test/domain/models"
	"go-gorm-test/usecase"
	"go-gorm-test/util"
	"net/http"

	"github.com/labstack/echo"
)

type FileController interface {
	FileGetAll(c echo.Context) (err error)
	FileUpload(c echo.Context) (err error)
	FileDownload(c echo.Context) (err error)
}

type fileController struct {
	Cfu usecase.FileUseCase
}

func NewFileController(cfu usecase.FileUseCase) FileController {
	return fileController{cfu}
}

type (
	filedownload struct {
		Filename string `json:"filename" validate:"required"`
	}
)

func (fc fileController) FileGetAll(c echo.Context) (err error) {
	post, statuscode, err := fc.Cfu.FileGetAllUseCase()
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (fc fileController) FileUpload(c echo.Context) (err error) {
	file, err := c.FormFile("file")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	post, statuscode, err := fc.Cfu.FileUploadUseCase(file)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.JSON(http.StatusOK, post)
}

func (fc fileController) FileDownload(c echo.Context) (err error) {
	fd := new(filedownload)
	err = util.BindValidate(c, fd)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, util.ErrorNoParameter)
	}
	post, statuscode, err := fc.Cfu.FileDownloadUseCase(fd.Filename)
	if err != nil {
		message := models.Message{
			Message: err.Error(),
		}
		return echo.NewHTTPError(statuscode, message)
	}
	return c.Blob(http.StatusOK, "image/jpeg", post)
}
