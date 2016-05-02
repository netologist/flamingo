package main

import (
	"github.com/hasanozgan/vodka"
	"github.com/hasanozgan/vodka/examples/basic/routes"
)

func main() {
	engine := vodka.Default()
	routes.Register(engine)
	engine.Run(":8000")
}
