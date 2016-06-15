package main

import (
	"github.com/hasanozgan/flamingo"
	"github.com/hasanozgan/flamingo/examples/basic/routes"
)

func main() {
	engine := flamingo.Default()
	routes.Register(engine)
	engine.Run(":8000")
}
