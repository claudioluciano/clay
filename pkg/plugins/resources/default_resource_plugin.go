package resources

import (
	"embed"
	"github.com/leap-fish/clay/pkg/clay"
	defaults2 "github.com/leap-fish/clay/pkg/plugins/resources/defaults"
	"github.com/leap-fish/clay/pkg/resource"
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
	resource.RegisterHandler("image", ".png", &defaults2.PngDefaultHandler{})
	resource.RegisterHandler("font", ".ttf", &defaults2.TtfDefaultHandler{})
	resource.RegisterHandler("sfx", ".ogg", &defaults2.OggDefaultHandler{})
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
