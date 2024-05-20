package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/dpi"
	"github.com/leap-fish/clay/components/spatial"
	"github.com/leap-fish/clay/components/sprite"
	"github.com/leap-fish/clay/render"
	"github.com/leap-fish/clay/resource"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"image"
)

type DefaultImageSystem struct {
	imageQuery *donburi.Query
}

func NewDefaultImageSystem() *DefaultImageSystem {
	return &DefaultImageSystem{
		imageQuery: donburi.NewQuery(
			filter.Contains(sprite.Component, spatial.TransformComponent)),
	}
}

func (s *DefaultImageSystem) Render(rg *render.RenderGraph, w donburi.World) {
	scaleFactor := dpi.GetScaleFactor(w)
	s.imageQuery.Each(w, func(entry *donburi.Entry) {
		spr := sprite.Component.Get(entry)
		tf := spatial.TransformComponent.Get(entry)
		if spr.Source == nil {
			res := resource.Get[image.Image](spr.Path)
			spr.Source = ebiten.NewImageFromImage(res)
		}

		rg.Add(func(world donburi.World, img *ebiten.Image, cam *camera.Camera) {
			render.Draw(spr.Source, render.ModeWorld, tf.Index).
				Scale(tf.Scale*scaleFactor).
				Origin(spr.Origin.XY()).
				Rotation(tf.Rotation).
				Filter(ebiten.FilterLinear).
				Position(tf.Position.XY()).
				Draw(img, cam)
		}, tf.Index)
	})
}
