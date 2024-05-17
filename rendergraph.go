package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/util/ecsutil"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"sort"
)

type DrawFunction func(world donburi.World, screen *ebiten.Image, camera *camera.Camera)

type phaseItem struct {
	drawFunction DrawFunction
	ordering     int
}

type RenderGraph struct {
	queue []phaseItem
}

func (rg *RenderGraph) Add(draw DrawFunction, order int) {
	rg.queue = append(rg.queue, phaseItem{
		drawFunction: draw,
		ordering:     order,
	})
}

// Prepare ensures the render graph's entries are sorted in the correct order.
func (rg *RenderGraph) Prepare() {
	// Sorts the queue according to ordering of the item.
	sort.Slice(rg.queue, func(i, j int) bool {
		return rg.queue[i].ordering < rg.queue[j].ordering
	})
}

func (rg *RenderGraph) Render(surface *ebiten.Image, w donburi.World) {
	surface.Clear()

	// Skips rendering if there's no queued operations.
	if len(rg.queue) == 0 {
		ebitenutil.DebugPrintAt(surface, "RENDERING SKIPPED\nNo items queued", 10, 10)
		return
	}

	// Find a camera, or early exit from render if there's no camera.
	// Displays a message to the consumer telling them to add the camera to the world.
	gameCamera := ecsutil.FirstOf[camera.Camera](camera.Component, w)
	if gameCamera == nil {
		log.Error("Cannot render because there is no camera present")
		ebitenutil.DebugPrintAt(surface, "NO CAMERA IS PRESENT\n> Add a camera to the ECS world.", 10, 10)
		return
	}

	// For every item in the queue, we can run the draw function.
	for _, item := range rg.queue {
		item.drawFunction(w, surface, gameCamera)
	}

	// Clear the queue
	clear(rg.queue)
}
