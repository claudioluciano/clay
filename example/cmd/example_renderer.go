package cmd

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/resource"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/debug"
	"image/color"
)

type ExampleRenderer struct {
	font     *text.GoTextFaceSource
	fontFace *text.GoTextFace
}

func (e *ExampleRenderer) Init(w donburi.World) {
	e.font = resource.Get[*text.GoTextFaceSource]("font:BaiJamjuree-Regular")
	e.fontFace = &text.GoTextFace{
		Source: e.font,
		Size:   14,
	}
}

func (e *ExampleRenderer) Render(rg *clay.RenderGraph, w donburi.World) {
	rg.Add(func(world donburi.World, screen *ebiten.Image, cam *camera.Camera) {
		var out bytes.Buffer
		for _, c := range debug.GetEntityCounts(world) {
			out.WriteString(fmt.Sprintf("> %s\n", c.String()))
		}

		out.WriteString(fmt.Sprintf("%#v", cam))

		op := &text.DrawOptions{}
		op.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
		op.GeoM.Translate(32, 32)
		op.LineSpacing = 14
		text.Draw(
			screen,
			fmt.Sprintf("Clay Editor\n%s", out.String()),
			e.fontFace,
			op,
		)
	}, 1)
}
