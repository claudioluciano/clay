package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay"
	"time"
)

type DefaultRendererModule struct {
}

func (d *DefaultRendererModule) Build(core *clay.Core) {
}

func (d *DefaultRendererModule) Ready(core *clay.Core) {
}

func (d *DefaultRendererModule) Update(dt time.Duration) {
}

func (d *DefaultRendererModule) Draw(screen *ebiten.Image) {
}
