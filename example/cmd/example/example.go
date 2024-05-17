package main

import (
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/example"
)

func main() {
	c := clay.New()
	c.Module(&example.ExampleModule{})

	c.Build()
	c.Run()
}
