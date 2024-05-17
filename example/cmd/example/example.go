package main

import (
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/example"
)

func main() {
	c := clay.New()
	c.Plugin(&example.ExamplePlugin{})

	c.Build()
	c.Run()
}
