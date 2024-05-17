package render

import (
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/components/camera"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

type DefaultRendererPlugin struct {
	ecs *ecs.ECS

	cam         *camera.Camera
	spriteQuery *donburi.Query
}

func (d *DefaultRendererPlugin) Build(core *clay.Core) {
	core.SubSystem(NewDefaultImageSystem())
}

func (d *DefaultRendererPlugin) Ready(core *clay.Core) {
}
