package resources

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/modules/resources/defaults"
	"github.com/leap-fish/clay/resource"
	log "github.com/sirupsen/logrus"
	"time"
)

type DefaultResourcesModule struct{}

func (r *DefaultResourcesModule) Update(dt time.Duration) {
}

func (r *DefaultResourcesModule) Draw(screen *ebiten.Image) {
}

func (r *DefaultResourcesModule) Ready(core *clay.Core) {}

func (r *DefaultResourcesModule) Build(core *clay.Core) {
	log.Info("Registering default handlers")
	resource.RegisterHandler("image", ".png", &defaults.PngDefaultHandler{})
	resource.RegisterHandler("font", ".ttf", &defaults.TtfDefaultHandler{})
}
