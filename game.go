package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"math"
	"reflect"
)

type ClayGame struct {
	ScrW, ScrH  int
	RenderScale float64
	Core        *Core
	World       donburi.World

	DrawSurface *ebiten.Image
}

func (g *ClayGame) Update() error {
	/*	for _, m := range g.Core.SortedPlugins {
		dt := time.Second / time.Duration(ebiten.TPS())
		m.Update(dt)
	}*/

	return nil
}

func (g *ClayGame) Draw(screen *ebiten.Image) {
	if g.DrawSurface == nil {
		g.DrawSurface = ebiten.NewImage(g.ScrW, g.ScrH)
	}

	bounds := g.DrawSurface.Bounds()
	if bounds.Dx() != g.ScrW || bounds.Dy() != g.ScrH {
		g.DrawSurface = ebiten.NewImage(g.ScrW, g.ScrH)
	}

	g.DrawSurface.Clear()
	for _, renderable := range g.Core.SubSystemRegistry.Renderables {
		log.Info("Queue ", reflect.TypeOf(renderable), renderable.Render)
		g.Core.RenderGraph.Add(renderable.Render, 0)
	}

	// Finds the graph:
	g.Core.RenderGraph.Prepare()

	g.Core.RenderGraph.Render(screen, g.DrawSurface, g.World)
}

func (g *ClayGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.ScrW = int(math.Round(float64(outsideWidth) / g.RenderScale))
	g.ScrH = int(math.Round(float64(outsideHeight) / g.RenderScale))
	return g.ScrW, g.ScrH
}
