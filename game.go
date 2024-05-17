package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/render"
	"math"
	"time"
)

type ClayGame struct {
	ScrW, ScrH  int
	RenderScale float64
	Core        *Core
}

func (g *ClayGame) Update() error {
	for _, m := range g.Core.SortedModules {
		dt := time.Second / time.Duration(ebiten.TPS())
		m.Update(dt)
	}

	return nil
}

func (g *ClayGame) Draw(screen *ebiten.Image) {
	for _, m := range g.Core.SortedModules {
		//log.Trace("Drawing for", reflect.TypeOf(m))
		m.Draw(screen)
	}

	cameraEntry, ok := camera.Query.First(g.Core.ECS.World)
	if ok && cameraEntry != nil {
		cam := camera.Component.Get(cameraEntry)
		render.Render(screen, cam)
	} else {
		g.Core.ECS.World.Create(camera.Component)
	}
}

func (g *ClayGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	w := math.Round(float64(outsideWidth) / g.RenderScale)
	h := math.Round(float64(outsideHeight) / g.RenderScale)
	return int(w), int(h)
}
