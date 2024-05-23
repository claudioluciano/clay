package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/dpi"
	"github.com/leap-fish/clay/pkg/components/spatial"
	"github.com/leap-fish/clay/pkg/components/sprite"
	r "github.com/leap-fish/clay/pkg/render"
	"github.com/leap-fish/clay/pkg/resource"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"image"
	"time"
)

type DefaultImageSystem struct {
	imageQuery *donburi.OrderedQuery[spatial.Transform]
}

func (s *DefaultImageSystem) Update(w donburi.World, dt time.Duration) {
	s.imageQuery.Each(w, func(entry *donburi.Entry) {
		tf := spatial.TransformComponent.Get(entry)
		tf.Index = int(tf.Position.Y)
	})
}

func NewDefaultImageSystem() *DefaultImageSystem {
	return &DefaultImageSystem{
		imageQuery: donburi.NewOrderedQuery[spatial.Transform](
			filter.Contains(sprite.Component, spatial.TransformComponent)),
	}
}

func (s *DefaultImageSystem) Render(rg *r.RenderGraph, w donburi.World) {
	scaleFactor := dpi.GetScaleFactor(w)
	s.imageQuery.EachOrdered(w, spatial.TransformComponent, func(entry *donburi.Entry) {
		spr := sprite.Component.Get(entry)
		tf := spatial.TransformComponent.Get(entry)
		if spr.Source == nil {
			res := resource.Get[image.Image](spr.Path)
			spr.Source = ebiten.NewImageFromImage(res)
		}

		rg.Add(func(world donburi.World, img *ebiten.Image, cam *camera.Camera) {
			r.Draw(spr.Source, r.ModeWorld, tf.Order()).
				Scale(tf.Scale*scaleFactor).
				Origin(spr.Origin.XY()).
				Rotation(tf.Rotation).
				Filter(ebiten.FilterLinear).
				Position(tf.Position.XY()).
				Draw(img, cam)
		}, tf.Order())
	})
}
