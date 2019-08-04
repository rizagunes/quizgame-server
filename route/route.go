package route

import (
	"github.com/labstack/echo"
	"quizgame-server/api"
	"quizgame-server/api/middlewares"
)

func New() *echo.Echo {
	e := echo.New()

	// Groups
	gameV1Group := e.Group("/game/v1")

	// Set middleware
	middlewares.SetJwtMiddleware(gameV1Group)

	api.MainGroup(e)
	api.GameV1Group(gameV1Group)

	return e
}
