package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/hasanozgan/flamingo"
	"github.com/hasanozgan/flamingo/examples/basic/context"
	"github.com/hasanozgan/flamingo/examples/basic/graphql"
	"github.com/hasanozgan/flamingo/examples/basic/handlers"
)

var (
	user *handlers.UserHandler
)

func init() {
	user = &handlers.UserHandler{}
}

func Register(engine *gin.Engine) {
	engine.GET("/", user.Index)
	engine.GET("/graphql", flamingo.NewGraphIQLHandler())
	engine.POST("/graphql", flamingo.NewGraphQLHandler(flamingo.GraphQLHandlerConfig{
		Schema:         graphql.NewGraphQLSchema(),
		Pretty:         true,
		ContextFactory: context.NewAppContext,
	}))
}
