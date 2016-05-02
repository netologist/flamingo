package vodka

import (
	"github.com/gin-gonic/gin"
)

type AppContext struct {
	ctx               *gin.Context
	dbClientFactory   *DbClientFactory
	httpClientFactory *HttpClientFactory
	loggerFactory     *LoggerFactory
}

func NewAppContext(ctx *gin.Context) *AppContext {
	loggerFactory := NewLoggerFactory(ctx.Keys)
	dbClientFactory := NewDbClientFactory(loggerFactory)
	httpClientFactory := NewHttpClientFactory(loggerFactory)

	return &AppContext{ctx, dbClientFactory, httpClientFactory, loggerFactory}

}

func (a *AppContext) IsLoggedIn() bool {
	return false
}

func (a *AppContext) UserInfo() *UserInfo {
	return &UserInfo{}
}

func (a *AppContext) Services() *Services {
	return &Services{a.dbClientFactory, a.httpClientFactory, a.loggerFactory}
}

func (c *AppContext) NewLogger(name string) (l *Logger) {
	return &Logger{}
}
