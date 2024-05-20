package plugins

import (
	"embed"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/plugins/audio"
	"github.com/leap-fish/clay/plugins/render"
	"github.com/leap-fish/clay/plugins/resources"
)

// DefaultPlugins adds core engine plugins that are required for rendering.
// You can override this by simply adding them manually to your core.Plugin() command.
// DefaultPlugins includes the resource loading system, camera system and rendering + audio systems.
// This provides some basics so you don't need to reinvent these unless your needs exceed that of the default
// implementation.
func DefaultPlugins(fs embed.FS) []clay.Plugin {
	return []clay.Plugin{
		resources.NewDefaultResourcesPlugin("assets", fs),
		&DefaultCameraPlugin{},
		&audio.DefaultAudioPlugin{},
		&render.DefaultRendererPlugin{},
	}
}
