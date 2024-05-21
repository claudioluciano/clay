package plugins

import (
	"github.com/leap-fish/clay/pkg/bundle"
	"github.com/leap-fish/clay/pkg/clay"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/events"
	"github.com/leap-fish/clay/pkg/util/ecsutil"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
)

type DefaultCameraPlugin struct{}

func (d *DefaultCameraPlugin) Build(core *clay.Core) {
}

func (d *DefaultCameraPlugin) Ready(core *clay.Core) {
	core.World.Create(camera.Component)
	events.EngineWindowSizeUpdated.Subscribe(core.World,
		func(w donburi.World, event events.WindowSizeUpdate) {
			cam := ecsutil.FirstOf[camera.Camera](camera.Component, w)
			if cam == nil {
				cameraBundle := bundle.New().
					With(camera.Component, camera.NewCamera(event.Width, event.Height, math.NewVec2(0, 0), 1.0))
				cameraBundle.Spawn(w)
			}
			cam.Resize(event.Width, event.Height)
			log.Info("Camera resized: ", event)
		},
	)
}
