package controllers

import (
	"go-gorm-test/domain/models"
	"net/http"

	"github.com/labstack/echo"
)

// Root return JSON
func Root(c echo.Context) error {
	post := &models.Message{
		Message: "this api is working!",
	}
	return c.JSON(http.StatusOK, post)
}
