package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	txt "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/dpi"
	"github.com/leap-fish/clay/pkg/components/spatial"
	"github.com/leap-fish/clay/pkg/components/text"
	"github.com/leap-fish/clay/pkg/render"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
)

type DefaultTextSystem struct {
	textQuery *donburi.OrderedQuery[spatial.Transform]

	op *txt.DrawOptions
}

func NewDefaultTextSystem() *DefaultTextSystem {
	return &DefaultTextSystem{
		textQuery: donburi.NewOrderedQuery[spatial.Transform](
			filter.Contains(text.Component, spatial.TransformComponent)),
		op: &txt.DrawOptions{},
	}
}

func (s *DefaultTextSystem) Render(rg *render.RenderGraph, w donburi.World) {
	scaleFactor := dpi.GetScaleFactor(w)
	s.textQuery.EachOrdered(w, spatial.TransformComponent, func(entry *donburi.Entry) {
		t := text.Component.Get(entry)
		tf := spatial.TransformComponent.Get(entry)

		face := &txt.GoTextFace{
			Source: t.Source,
			Size:   t.Size * scaleFactor,
		}

		op := &txt.DrawOptions{}
		op.ColorScale.ScaleWithColor(t.Color)
		op.LineSpacing = (t.Size * scaleFactor) * t.LineHeight
		op.PrimaryAlign = t.PrimaryAlign
		op.SecondaryAlign = t.SecondaryAlign
		op.GeoM.Translate(tf.Position.XY())

		rg.Add(func(world donburi.World, img *ebiten.Image, cam *camera.Camera) {
			txt.Draw(img, t.Content.String(), face, op)
		}, tf.Index)
	})
}
