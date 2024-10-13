package render

import (
	"github.com/leap-fish/clay/pkg/clay"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/dpi"
	log "github.com/sirupsen/logrus"
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
		NewDefaultAnimSpriteSystem(),
		NewDefaultTextSystem(),
	)
}

func (d *DefaultRendererPlugin) Ready(core *clay.Core) {
	if clay.LaunchOptions.UseDPIScaling {
		log.Trace("Created DPI scaler component")
		core.World.Create(dpi.Component)
	}
}
