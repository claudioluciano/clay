package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/sprite"
	"github.com/leap-fish/clay/components/transform"
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
		imageQuery: donburi.NewQuery(filter.Contains(sprite.Component)),
	}
}

func (s DefaultImageSystem) Render(w donburi.World, img *ebiten.Image, cam *camera.Camera) {
	s.imageQuery.Each(w, func(entry *donburi.Entry) {
		spr := sprite.Component.Get(entry)
		tf := transform.Component.Get(entry)
		if spr.Source == nil {
			res := resource.Get[image.Image](spr.Path)
			spr.Source = ebiten.NewImageFromImage(res)
		}

		render.Draw(spr.Source, render.RenderModeCanvas, tf.Index).
			Scale(tf.Scale).
			OriginMul(spr.Origin.XY()).
			Rotation(tf.Rotation).
			Position(tf.Position.XY()).Draw(img, cam)
	})
}
