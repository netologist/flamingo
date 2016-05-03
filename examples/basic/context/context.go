package context

import (
	"github.com/gin-gonic/gin"
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/services"
)

type AppContext struct {
	ctx               *gin.Context
	dbClientFactory   *vodka.DbClientFactory
	httpClientFactory *vodka.HttpClientFactory
	loggerFactory     *vodka.LoggerFactory
}

func NewAppContext(ctx *gin.Context) *AppContext {
	fields := map[string]interface{}{}
	fields["X-Correlation-ID"] = ctx.Request.Header["X-Correlation-ID"]

	loggerFactory := vodka.NewLoggerFactory(fields)
	dbClientFactory := vodka.NewDbClientFactory(loggerFactory)
	httpClientFactory := vodka.NewHttpClientFactory(loggerFactory)

	return &AppContext{ctx, dbClientFactory, httpClientFactory, loggerFactory}
}

func (a *AppContext) IsLoggedIn() bool {
	return false
}

func (a *AppContext) UserInfo() *vodka.UserInfo {
	return &vodka.UserInfo{}
}

func (a *AppContext) Services() *services.Services {
	return services.New(a.dbClientFactory, a.httpClientFactory, a.loggerFactory)
}

func (c *AppContext) NewLogger(name string) (l *vodka.Logger) {
	return &vodka.Logger{}
}
