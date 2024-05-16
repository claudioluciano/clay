package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"math"
	"time"
)

type ClayGame struct {
	ScrW, ScrH  int
	RenderScale float64
	Core        *Core
}

func (g *ClayGame) Update() error {
	for _, m := range g.Core.Modules {
		dt := time.Second / time.Duration(ebiten.TPS())
		m.Update(dt)
	}

	return nil
}

func (g *ClayGame) Draw(screen *ebiten.Image) {
	for _, m := range g.Core.Modules {
		m.Draw(screen)
	}
}

func (g *ClayGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w := math.Round(float64(outsideWidth) / g.RenderScale)
	h := math.Round(float64(outsideHeight) / g.RenderScale)
	return int(w), int(h)
}
