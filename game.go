package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/modules/render"
	"github.com/yohamta/donburi"
	"math"
	"time"
)

type ClayGame struct {
	ScrW, ScrH  int
	RenderScale float64
	Core        *Core
	World       donburi.World
}

func (g *ClayGame) Update() error {
	for _, m := range g.Core.SortedModules {
		dt := time.Second / time.Duration(ebiten.TPS())
		m.Update(dt)
	}

	return nil
}

func (g *ClayGame) Draw(screen *ebiten.Image) {
	/*	for _, m := range g.Core.SortedModules {
		moduleRenderable, isRenderable := m.(RenderableModule)
		if isRenderable {
			//render.QueueFunc(moduleRenderable.QueueDraw)
			log.Info("Is renderable", moduleRenderable)
			moduleRenderable.Draw(screen)
		}
		//log.Trace("Drawing for", reflect.TypeOf(m))
	}*/

	// Finds the graph:
	render.GraphQuery.First(g.World)

}

func (g *ClayGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w := math.Round(float64(outsideWidth) / g.RenderScale)
	h := math.Round(float64(outsideHeight) / g.RenderScale)
	return int(w), int(h)
}
