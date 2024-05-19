package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/dpi"
	ev "github.com/leap-fish/clay/events"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"math"
	"time"
)

type ClayGame struct {
	ScrW, ScrH int
	Core       *Core
	World      donburi.World
}

func (g *ClayGame) Init() {
	for _, initializable := range g.Core.SubSystemRegistry.Initializables {
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
	scaleFactor := dpi.GetScaleFactor(g.Core.World)
	renderScale := g.Core.Options.RenderScale
	newW := int(math.Round(float64(outsideWidth)/renderScale) * scaleFactor)
	newH := int(math.Round(float64(outsideHeight)/renderScale) * scaleFactor)

	if newW != g.ScrW || newH != g.ScrH {
		ev.EngineWindowSizeUpdated.Publish(g.World, ev.WindowSizeUpdate{
			Width:  newW,
			Height: newH,
		})
		g.ScrW = newW
		g.ScrH = newH
		log.
			WithField("windowWidth", g.ScrW).
			WithField("windowHeight", g.ScrH).
			WithField("scale", renderScale).
			WithField("scaleFactor", scaleFactor).
			Trace("Layout size has changed")
	}

	return g.ScrW, g.ScrH
}
