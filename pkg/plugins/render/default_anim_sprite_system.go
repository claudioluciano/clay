package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/pkg/components/animsprite"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/dpi"
	"github.com/leap-fish/clay/pkg/components/spatial"
	r "github.com/leap-fish/clay/pkg/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"github.com/yohamta/ganim8/v2"
	"time"
)

type DefaultAnimSpriteSystem struct {
	imageQuery *donburi.OrderedQuery[spatial.Transform]
}

func (s *DefaultAnimSpriteSystem) Update(w donburi.World, dt time.Duration) {
	s.imageQuery.Each(w, func(entry *donburi.Entry) {
		tf := spatial.TransformComponent.Get(entry)
		tf.Index = int(tf.Position.Y)
	})

	s.imageQuery.EachOrdered(w, spatial.TransformComponent, func(entry *donburi.Entry) {
		spr := animsprite.Component.Get(entry)

		animation := spr.Animations[spr.CurrentAnimation]
		if animation != nil {
			animation.InternalAnimation.Update()
		}
	})
}

func NewDefaultAnimSpriteSystem() *DefaultAnimSpriteSystem {
	return &DefaultAnimSpriteSystem{
		imageQuery: donburi.NewOrderedQuery[spatial.Transform](
			filter.Contains(animsprite.Component, spatial.TransformComponent)),
	}
}

func (s *DefaultAnimSpriteSystem) Render(rg *r.RenderGraph, w donburi.World) {
	scaleFactor := dpi.GetScaleFactor(w)
	s.imageQuery.EachOrdered(w, spatial.TransformComponent, func(entry *donburi.Entry) {
		spr := animsprite.Component.Get(entry)
		tf := spatial.TransformComponent.Get(entry)

		animation := spr.Animations[spr.CurrentAnimation]
		if animation != nil {
			rg.Add(func(world donburi.World, img *ebiten.Image, cam *camera.Camera) {
				spr.Image.Clear()
				animation.InternalAnimation.Draw(spr.Image, ganim8.DrawOpts(0, 0, 0, 1, 1, 0, 0))
				r.Draw(spr.Image, r.ModeWorld, tf.Order()).
					Scale(tf.Scale*scaleFactor).
					Origin(spr.Origin.XY()).
					Rotation(tf.Rotation).
					Filter(spr.Filter).
					Position(tf.Position.XY()).
					Draw(img, cam)
			}, tf.Order())
		}
	})
}
