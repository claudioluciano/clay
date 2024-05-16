package main

import (
	"clay"
	"clay/editor"
)

func main() {
	c := clay.New()
	c.Module(&editor.EditorModule{})
	c.Run()
}
