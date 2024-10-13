package sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/pkg/resource"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"image"
)

var Component = donburi.NewComponentType[Sprite]()

type Sprite struct {
	Path resource.Path

	FlipX bool
	FlipY bool

	Origin math.Vec2

	CustomSize *math.Vec2
	CustomRect *image.Rectangle

	Filter ebiten.Filter
	Source *ebiten.Image
}
