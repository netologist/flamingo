package vodka

import "github.com/gin-gonic/gin"

//type ContextFactory func(g *gin.Context)
type HandlerFunc func(c Context)

type Engine struct {
	engine  *gin.Engine
	context Context
}

func (e *Engine) Run(addr string) {
	e.engine.Run(addr)
}

func (c *Engine) RegisterContext(context Context) {
	c.context = context
}

func (e *Engine) GET(relativePath string, handler HandlerFunc) gin.IRoutes {
	return e.Handle("GET", relativePath, handler)
}

func (e *Engine) Handle(method string, relativePath string, handler HandlerFunc) gin.IRoutes {
	return e.engine.Handle(method, relativePath, func(c *gin.Context) {
		//e.context
		//appContext := e.contextFactory(c).(Context)
		handler(e.context)
	})
}
