package main

import (
	"quizgame-server/route"
)

func main() {
	e := route.New()
	e.Start(":8000")
}
