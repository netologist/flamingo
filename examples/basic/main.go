package main

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/context"
	"github.com/hasanozgan/vodka/examples/basic/routes"
)

func main() {
	engine := vodka.Default()
	engine.RegisterContext(context.AppContext{})
	routes.Register(engine)
	engine.Run(":8000")
}
