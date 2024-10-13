package animsprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/pkg/resource"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/ganim8/v2"
	"image"
)

var Component = donburi.NewComponentType[AnimSprite]()

type Animation struct {
	Frames    []any
	Durations any

	InternalAnimation *ganim8.Animation
}

func NewAnimation(durations any, frames ...any) *Animation {
	return &Animation{
		Frames:            frames,
		Durations:         durations,
		InternalAnimation: nil,
	}
}

type AnimSprite struct {
	SpritesheetPath resource.Path
	Source          *ebiten.Image
	Image           *ebiten.Image

	Grid             *ganim8.Grid
	Animations       map[string]*Animation
	CurrentAnimation string

	FlipX bool
	FlipY bool

	Origin math.Vec2
	Filter ebiten.Filter
}

type SpriteSheetSize struct {
	FrameHeight int
	FrameWidth  int
	ImageHeight int
	ImageWidth  int
}

func New(path resource.Path, animation string, spriteSheet SpriteSheetSize, animations map[string]*Animation, filter ebiten.Filter) AnimSprite {
	res := resource.Get[image.Image](path)
	source := ebiten.NewImageFromImage(res)

	grid := ganim8.NewGrid(spriteSheet.FrameWidth, spriteSheet.FrameHeight, spriteSheet.ImageWidth, spriteSheet.ImageHeight)

	for _, animation := range animations {
		animation.InternalAnimation = ganim8.New(source, grid.Frames(animation.Frames...), animation.Durations)
	}

	return AnimSprite{
		Image:            ebiten.NewImage(spriteSheet.FrameWidth, spriteSheet.FrameHeight),
		Source:           source,
		SpritesheetPath:  path,
		Grid:             grid,
		CurrentAnimation: animation,
		Animations:       animations,
		Filter:           filter,
	}
}
