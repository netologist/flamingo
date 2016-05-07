package graphql

import (
	"github.com/hasanozgan/vodka/examples/basic/context"

	"github.com/graphql-go/graphql"
)

func NewGraphQLSchema() (schema graphql.Schema) {
	// Schema
	fields := graphql.Fields{
		"hello": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {

				ctx := p.Context.(context.AppContext)
				logger := ctx.NewLogger("User.Index Handler")
				logger.Info("hede hodo")

				return "world", nil
			},
		},
	}
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: fields}
	schemaConfig := graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)}
	schema, _ = graphql.NewSchema(schemaConfig)
	return
}
