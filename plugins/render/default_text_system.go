package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	txt "github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/spatial"
	"github.com/leap-fish/clay/components/text"
	log "github.com/sirupsen/logrus"
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
	s.textQuery.Each(w, func(entry *donburi.Entry) {
		t := text.Component.Get(entry)
		tf := spatial.TransformComponent.Get(entry)

		// Skip non valid fontfaces
		if t.FontFace == nil {
			log.Info("Invalid fontface")
			return
		}

		op := &txt.DrawOptions{}
		op.ColorScale.ScaleWithColor(t.Color)
		op.LineSpacing = float64(t.Size) * t.LineHeight
		op.GeoM.Translate(tf.Position.XY())

		rg.Add(func(world donburi.World, img *ebiten.Image, cam *camera.Camera) {
			txt.Draw(img, t.Content.String(), t.FontFace, op)
		}, tf.Index)
	})
}
