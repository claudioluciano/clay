package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/dpi"
	"github.com/leap-fish/clay/config"
	ev "github.com/leap-fish/clay/events"
	"github.com/leap-fish/clay/render"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/events"
	"math"
	"time"
)

type ClayGame struct {
	ScrW, ScrH int

	//Core       *Core
	World donburi.World

	subSystems *SubSystemRegistry
	plugins    *PluginRegistry

	// Used for rendering
	RenderGraph *render.RenderGraph

	Options *config.LaunchOptions
}

func NewClayGame(w donburi.World, subSystems *SubSystemRegistry, plugins *PluginRegistry, options *config.LaunchOptions) *ClayGame {
	return &ClayGame{
		ScrW:        0,
		ScrH:        0,
		World:       w,
		RenderGraph: &render.RenderGraph{},
		Options:     options,

		plugins:    plugins,
		subSystems: subSystems,
	}
}

func (g *ClayGame) Init() {
	// Defaults set by ebiten
	ebiten.SetWindowSize(g.Options.WindowWidth, g.Options.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(g.Options.VsyncMode)

	g.RenderGraph = &render.RenderGraph{}

	for _, initializable := range g.subSystems.Initializables {
		initializable.Init(g.World)
	}
}

func (g *ClayGame) Update() error {
	events.ProcessAllEvents(g.World)

	for _, updatable := range g.subSystems.Updatables {
		//g.Core.RenderGraph.Add(updatable.Render, i)
		updatable.Update(g.World, time.Second/time.Duration(ebiten.TPS()))
	}

	return nil
}

func (g *ClayGame) Draw(screen *ebiten.Image) {
	for _, renderable := range g.subSystems.Renderables {
		renderable.Render(g.RenderGraph, g.World)
	}

	g.RenderGraph.Prepare()
	g.RenderGraph.Render(screen, g.World)
}

func (g *ClayGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	dpiScaleFactor := dpi.GetScaleFactor(g.World)
	renderScale := g.Options.RenderScale
	newW := int(math.Round(float64(outsideWidth)/renderScale) * dpiScaleFactor)
	newH := int(math.Round(float64(outsideHeight)/renderScale) * dpiScaleFactor)

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
			WithField("renderScale", renderScale).
			WithField("dpiScaleFactor", dpiScaleFactor).
			Trace("Layout size has changed")
	}

	return g.ScrW, g.ScrH
}
