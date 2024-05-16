package main

import (
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/modules/editor"
	"github.com/leap-fish/clay/modules/resources"
)

func main() {
	c := clay.New()
	c.Module(&editor.EditorModule{}, &resources.DefaultResourcesModule{})
	c.Run()
}
