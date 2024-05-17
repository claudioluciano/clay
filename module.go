package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/components/camera"
	"time"
)

type Module interface {
	Build(core *Core)
	Ready(core *Core)
	Update(dt time.Duration)
}

type RenderableModule interface {
	Draw(screen *ebiten.Image, camera *camera.Camera)
}

type Sortable interface {
	Order() int
}
