package resources

import (
	"embed"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/plugins/resources/defaults"
	"github.com/leap-fish/clay/resource"
	log "github.com/sirupsen/logrus"
	"math"
)

type DefaultResourcesPlugin struct {
	FileSystem embed.FS
	Path       string
}

func NewDefaultResourcesPlugin(path string, fs embed.FS) *DefaultResourcesPlugin {
	return &DefaultResourcesPlugin{
		Path:       path,
		FileSystem: fs,
	}
}

func (r *DefaultResourcesPlugin) Order() int {
	return math.MinInt32
}

func (r *DefaultResourcesPlugin) Build(core *clay.Core) {
	log.Info("Registering default handlers")
	resource.RegisterHandler("image", ".png", &defaults.PngDefaultHandler{})
	resource.RegisterHandler("font", ".ttf", &defaults.TtfDefaultHandler{})
}

func (r *DefaultResourcesPlugin) Ready(core *clay.Core) {
	resourceErrs := resource.LoadFromEmbedFolder(r.Path, r.FileSystem)
	if len(resourceErrs) > 0 {
		log.
			WithField("errors", resourceErrs).
			WithField("path", r.Path).
			WithField("fs", r.FileSystem).
			Errorf("Unable to load %d files from embedded file system", len(resourceErrs))
	}
}
