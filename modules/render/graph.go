package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/camera"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/filter"
	"sort"
)

var GraphComponent = donburi.NewComponentType[Graph]()
var GraphQuery = donburi.NewQuery(filter.Contains(GraphComponent))

type phaseItem struct {
	entity       donburi.Entity
	drawFunction func(screen *ebiten.Image, camera *camera.Camera)
	ordering     int
}

type Graph struct {
	queue []phaseItem
}

// Prepare ensures the render graph's entries are sorted in the correct order.
func (g *Graph) Prepare() {
	// Sorts the queue according to ordering of the item.
	sort.Slice(g.queue, func(i, j int) bool {
		return g.queue[i].ordering < g.queue[j].ordering
	})
}

func (g *Graph) Render(surface *ebiten.Image, w donburi.World) {
	surface.Clear()

	cameraEntry, ok := camera.Query.First(w)
	if !ok || cameraEntry == nil {
		log.Error("Cannot render because there is no camera present")
		return
	}

	// Acquire the game camera component and use it when calling the queue drawing functions.
	gameCamera := camera.Component.Get(cameraEntry)
	for _, item := range g.queue {
		item.drawFunction(surface, gameCamera)
	}

	// Clear the queue
	clear(g.queue)
}
