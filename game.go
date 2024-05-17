package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	ev "github.com/leap-fish/clay/events"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"math"
	"time"
)

type ClayGame struct {
	ScrW, ScrH  int
	RenderScale float64
	Core        *Core
	World       donburi.World
}

func (g *ClayGame) Init() {
	for _, initializable := range g.Core.SubSystemRegistry.Initializables {
		//g.Core.RenderGraph.Add(initializable.Render, i)
		initializable.Init(g.World)
	}
}

func (g *ClayGame) Update() error {
	events.ProcessAllEvents(g.World)

	for _, updatable := range g.Core.SubSystemRegistry.Updatables {
		//g.Core.RenderGraph.Add(updatable.Render, i)
		updatable.Update(g.World, time.Second/time.Duration(ebiten.TPS()))
	}

	return nil
}

func (g *ClayGame) Draw(screen *ebiten.Image) {
	for _, renderable := range g.Core.SubSystemRegistry.Renderables {
		renderable.Render(g.Core.RenderGraph, g.World)
	}

	g.Core.RenderGraph.Prepare()
	g.Core.RenderGraph.Render(screen, g.World)
}

func (g *ClayGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	newW := int(math.Round(float64(outsideWidth) / g.RenderScale))
	newH := int(math.Round(float64(outsideHeight) / g.RenderScale))

	if newW != g.ScrW || newH != g.ScrH {
		ev.EngineWindowSizeUpdated.Publish(g.World, ev.WindowSizeUpdate{
			Width:  newW,
			Height: newH,
		})
		g.ScrW = newW
		g.ScrH = newH
		log.
			WithField("w", g.ScrW).
			WithField("h", g.ScrH).
			Trace("Layout size has changed")
	} else {
	}

	return g.ScrW, g.ScrH
}
