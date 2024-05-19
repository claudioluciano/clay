package render

import (
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/dpi"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type DefaultRendererPlugin struct {
	ecs *ecs.ECS

	cam         *camera.Camera
	spriteQuery *donburi.Query
}

func (d *DefaultRendererPlugin) Build(core *clay.Core) {
	core.SubSystem(
		NewDefaultImageSystem(),
		NewDefaultTextSystem(),
	)
}

func (d *DefaultRendererPlugin) Ready(core *clay.Core) {
	if core.Options.UseDPIScaling {
		core.World.Create(dpi.Component)
	}
}
