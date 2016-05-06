package context

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/services"
)

type AppContext struct {
	vodka.Context
	ctx               *gin.Context
	dbClientFactory   *vodka.DbClientFactory
	httpClientFactory *vodka.HttpClientFactory
	loggerFactory     *vodka.LoggerFactory
}

func (c AppContext) SetDefaultFields(ctx *gin.Context) map[string]interface{} {
	return map[string]interface{}{}
}

func (c AppContext) LoggerFactory() vodka.LoggerFactory {
	return vodka.LoggerFactory{}
}
func (c AppContext) HttpClientFactory() vodka.HttpClientFactory {
	return vodka.HttpClientFactory{}
}
func (c AppContext) DbClientFactory() vodka.DbClientFactory {
	return vodka.DbClientFactory{}
}

func NewAppContext(ctx *gin.Context) AppContext {
	fields := map[string]interface{}{}
	fields["X-Correlation-ID"] = ctx.Request.Header["X-Correlation-ID"]

	loggerFactory := vodka.NewLoggerFactoryWithFields(fields)
	dbClientFactory := vodka.NewDbClientFactory(loggerFactory)
	httpClientFactory := vodka.NewHttpClientFactory(loggerFactory)

	context := AppContext{ctx: ctx, dbClientFactory: dbClientFactory, httpClientFactory: httpClientFactory, loggerFactory: loggerFactory}
	return context
}

func (a *AppContext) IsLoggedIn() bool {
	return false
}

func (a *AppContext) UserInfo() *UserInfo {
	return &UserInfo{}
}

func (a *AppContext) Services() *services.Services {
	return services.New(a.dbClientFactory, a.httpClientFactory, a.loggerFactory)
}

func (a *AppContext) NewLogger(name string) vodka.Logger {
	return a.loggerFactory.NewLogger(name)
}
