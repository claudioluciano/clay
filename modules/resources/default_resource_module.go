package resources

import (
	"embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/modules/resources/defaults"
	"github.com/leap-fish/clay/resource"
	log "github.com/sirupsen/logrus"
	"math"
	"time"
)

type DefaultResourcesModule struct {
	FileSystem embed.FS
	Path       string
}

func NewDefaultResourcesModule(path string, fs embed.FS) *DefaultResourcesModule {
	return &DefaultResourcesModule{
		Path:       path,
		FileSystem: fs,
	}
}

func (r *DefaultResourcesModule) Order() int {
	return math.MinInt32
}

func (r *DefaultResourcesModule) Build(core *clay.Core) {
	log.Info("Registering default handlers")
	resource.RegisterHandler("image", ".png", &defaults.PngDefaultHandler{})
	resource.RegisterHandler("font", ".ttf", &defaults.TtfDefaultHandler{})
}

func (r *DefaultResourcesModule) Ready(core *clay.Core) {
	resourceErrs := resource.LoadFromEmbedFolder(r.Path, r.FileSystem)
	if len(resourceErrs) > 0 {
		log.
			WithField("errors", resourceErrs).
			WithField("path", r.Path).
			WithField("fs", r.FileSystem).
			Errorf("Unable to load %d files from embedded file system", len(resourceErrs))
	}
}

func (r *DefaultResourcesModule) Update(dt time.Duration) {
}

func (r *DefaultResourcesModule) Draw(screen *ebiten.Image) {
}
