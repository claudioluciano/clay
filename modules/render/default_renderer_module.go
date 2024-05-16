package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay"
	"github.com/yohamta/donburi/ecs"
	"time"
)

type DefaultRendererModule struct {
	ecs *ecs.ECS
}

func (d *DefaultRendererModule) Build(core *clay.Core) {
}

func (d *DefaultRendererModule) Ready(core *clay.Core) {
}

func (d *DefaultRendererModule) Update(dt time.Duration) {
}

func (d *DefaultRendererModule) Draw(screen *ebiten.Image) {
}
