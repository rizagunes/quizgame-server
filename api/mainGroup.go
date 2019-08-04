package api

import (

	"github.com/labstack/echo"
	"quizgame-server/api/handlers"
)

func MainGroup(e *echo.Echo) {

	e.GET("/", handlers.HomePage)
	e.POST("/login", handlers.Login)
}


