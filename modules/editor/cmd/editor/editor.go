package main

import (
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/modules/editor"
)

func main() {
	c := clay.New()
	c.Module(&editor.EditorModule{})

	c.Build()
	c.Run()
}
