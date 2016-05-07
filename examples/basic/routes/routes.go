package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/context"
	"github.com/hasanozgan/vodka/examples/basic/graphql"
	"github.com/hasanozgan/vodka/examples/basic/handlers"
)

var (
	user *handlers.UserHandler
)

func init() {
	user = &handlers.UserHandler{}
}

func Register(engine *gin.Engine) {
	engine.GET("/", user.Index)
	engine.GET("/graphql", vodka.NewGraphIQLHandler())
	engine.POST("/graphql", vodka.NewGraphQLHandler(vodka.GraphQLHandlerConfig{
		Schema:         graphql.NewGraphQLSchema(),
		Pretty:         true,
		ContextFactory: context.NewAppContext,
	}))
}
