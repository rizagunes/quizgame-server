package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func HomePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func SecretPage(c echo.Context) error {
	return c.String(http.StatusOK, "└[∵┌]└[ ∵ ]┘[┐∵]┘")
}

