package vodka

import "github.com/gin-gonic/gin"

func Default() *Engine {
	loadConfig()
	engine := &Engine{engine: gin.Default()}
	engine.RegisterContext(AppContext{})
	return engine
}
