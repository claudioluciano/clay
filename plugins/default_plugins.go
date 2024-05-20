package plugins

import (
	"embed"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/plugins/render"
	"github.com/leap-fish/clay/plugins/resources"
)

// DefaultPlugins adds
func DefaultPlugins(fs embed.FS) []clay.Plugin {
	return []clay.Plugin{
		resources.NewDefaultResourcesPlugin("assets", fs),
		&DefaultCameraPlugin{},
		&render.DefaultRendererPlugin{},
	}
}
