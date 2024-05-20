package example

import (
	"embed"
	"flag"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/bundle"
	"github.com/leap-fish/clay/components/spatial"
	"github.com/leap-fish/clay/components/sprite"
	"github.com/leap-fish/clay/example/cmd"
	"github.com/leap-fish/clay/plugins"
	log "github.com/sirupsen/logrus"
	m "github.com/yohamta/donburi/features/math"
	"math"
	"time"
)

var (
	windowWidthFlag  = flag.Int("width", 800, "window width")
	windowHeightFlag = flag.Int("height", 600, "window height")
)

//go:embed assets
var EditorFiles embed.FS

type ExamplePlugin struct {
}

var imageSprite = bundle.New().
	With(spatial.TransformComponent, spatial.Transform{}).
	With(sprite.Component, sprite.Sprite{
		Path: "image:image",
	})

func (e *ExamplePlugin) Update(dt time.Duration) {
}

func (e *ExamplePlugin) Ready(core *clay.Core) {
	ent := imageSprite.Spawn(core.World)
	imageEntry := core.World.Entry(ent)
	spatial.TransformComponent.Set(imageEntry, &spatial.Transform{
		Position: m.Vec2{},
		Scale:    0.3,
	})

	ent2 := imageSprite.Spawn(core.World)
	imageEntry2 := core.World.Entry(ent2)
	spatial.TransformComponent.Set(imageEntry2, &spatial.Transform{
		Index:    3,
		Position: m.NewVec2(-100, -100),
		Rotation: 90 * (math.Pi / 180),
		Scale:    0.1,
	})
}

func (e *ExamplePlugin) Build(core *clay.Core) {
	core.Plugin(
		plugins.DefaultPlugins(EditorFiles)...,
	)

	core.SubSystem(&cmd.ExampleSystem{})

	core.LaunchOptions(clay.LaunchOptions{
		WindowWidth:   *windowWidthFlag,
		WindowHeight:  *windowHeightFlag,
		UseDPIScaling: true,
		RenderScale:   1.0,
	})
	log.Tracef("Window Size set to %dx%d", *windowWidthFlag, *windowHeightFlag)
}
