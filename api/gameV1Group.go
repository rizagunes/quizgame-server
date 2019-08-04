package api

import (

	"github.com/labstack/echo"
	"quizgame-server/api/handlers"
)

func GameV1Group(e *echo.Group) {

	e.GET("/", handlers.SecretPage)

}


