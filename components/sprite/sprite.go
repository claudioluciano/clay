package sprite

import (
	"github.com/leap-fish/clay/resource"
	"github.com/yohamta/donburi"
)

var Component = donburi.NewComponentType[Sprite]()

type Sprite struct {
	Path resource.Path
}
