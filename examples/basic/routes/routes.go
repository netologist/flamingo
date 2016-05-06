package routes

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/handlers"
)

var (
	userHandler *handlers.UserHandler
)

func init() {
	userHandler = &handlers.UserHandler{}
}

func Register(engine *vodka.Engine) {
	engine.GET("/", userHandler.Index)
}
