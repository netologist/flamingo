package vodka

import (
	"github.com/gin-gonic/gin"
)

type Context interface {
	SetGinContext(ctx *gin.Context)
	//SetDefaultFields(ctx *gin.Context) map[string]interface{}
	LoggerFactory() LoggerFactory
	HttpClientFactory() HttpClientFactory
	DbClientFactory() DbClientFactory
}

type AppContext struct {
	Context
	ctx               *gin.Context
	dbClientFactory   *DbClientFactory
	httpClientFactory *HttpClientFactory
	loggerFactory     *LoggerFactory
}

func (c AppContext) SetGinContext(ctx *gin.Context) {
	c.ctx = ctx
}

func (c AppContext) SetDefaultFields(ctx *gin.Context) map[string]interface{} {
	return map[string]interface{}{}
}

func (c AppContext) LoggerFactory() LoggerFactory {
	return LoggerFactory{}
}
func (c AppContext) HttpClientFactory() HttpClientFactory {
	return HttpClientFactory{}
}
func (c AppContext) DbClientFactory() DbClientFactory {
	return DbClientFactory{}
}

func NewAppContext(ctx *gin.Context,
	dbClientFactory *DbClientFactory,
	httpClientFactory *HttpClientFactory,
	loggerFactory *LoggerFactory) AppContext {
	return AppContext{ctx: ctx, dbClientFactory: dbClientFactory, httpClientFactory: httpClientFactory, loggerFactory: loggerFactory}
}

func DefaultContextFactory(ctx *gin.Context) interface{} {
	loggerFactory := NewLoggerFactory()
	dbClientFactory := NewDbClientFactory(loggerFactory)
	httpClientFactory := NewHttpClientFactory(loggerFactory)

	context := NewAppContext(ctx, dbClientFactory, httpClientFactory, loggerFactory)
	context.SetDefaultFields(ctx)
	return context
}
