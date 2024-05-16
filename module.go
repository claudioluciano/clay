package clay

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

type Module interface {
	Build(core *Core)
	Ready(core *Core)
	Update(dt time.Duration)
	Draw(screen *ebiten.Image)
}

type Sortable interface {
	Order() int
}
