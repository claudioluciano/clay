package main

import (
	"_example"
	"github.com/leap-fish/clay"
)

func main() {
	c := clay.New()
	c.Plugin(&_example.ExamplePlugin{})
	c.Run()
}
