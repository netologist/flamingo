package flamingo

import (
	"github.com/gin-gonic/gin"
)

func Default() *gin.Engine {
	loadConfig()

	app := gin.Default()
	return app
}
