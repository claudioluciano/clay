package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	txt "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/spatial"
	"github.com/leap-fish/clay/components/text"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type DefaultTextSystem struct {
	textQuery *donburi.Query

	op *txt.DrawOptions
}

func NewDefaultTextSystem() *DefaultTextSystem {
	return &DefaultTextSystem{
		textQuery: donburi.NewQuery(
			filter.Contains(text.Component, spatial.TransformComponent)),
		op: &txt.DrawOptions{},
	}
}

func (s *DefaultTextSystem) Render(rg *clay.RenderGraph, w donburi.World) {
	scaleFactor := ebiten.Monitor().DeviceScaleFactor()

	s.textQuery.Each(w, func(entry *donburi.Entry) {
		t := text.Component.Get(entry)
		tf := spatial.TransformComponent.Get(entry)

		face := &txt.GoTextFace{
			Source: t.Source,
			Size:   t.Size * scaleFactor,
		}

		op := &txt.DrawOptions{}
		op.ColorScale.ScaleWithColor(t.Color)
		op.LineSpacing = (t.Size * scaleFactor) * t.LineHeight
		op.GeoM.Translate(tf.Position.XY())

		rg.Add(func(world donburi.World, img *ebiten.Image, cam *camera.Camera) {
			txt.Draw(img, t.Content.String(), face, op)
		}, tf.Index)
	})
}
