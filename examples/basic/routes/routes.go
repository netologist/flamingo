package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanozgan/vodka/examples/basic/handlers"
)

var (
	userHandler *handlers.UserHandler
)

func init() {
	userHandler = &handlers.UserHandler{}
}

func Register(engine *gin.Engine) {
	engine.GET("/", userHandler.Index)
}
