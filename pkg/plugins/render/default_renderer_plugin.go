package render

import (
	"github.com/leap-fish/clay/pkg"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/dpi"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type DefaultRendererPlugin struct {
	ecs *ecs.ECS

	cam         *camera.Camera
	spriteQuery *donburi.Query
}

func (d *DefaultRendererPlugin) Build(core *pkg.Core) {
	core.SubSystem(
		NewDefaultImageSystem(),
		NewDefaultTextSystem(),
	)
}

func (d *DefaultRendererPlugin) Ready(core *pkg.Core) {
	if core.Options.UseDPIScaling {
		core.World.Create(dpi.Component)
	}
}
