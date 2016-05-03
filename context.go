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
