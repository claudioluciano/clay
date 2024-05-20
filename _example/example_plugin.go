package _example

import (
	"embed"
	"github.com/leap-fish/clay/pkg"
	"github.com/leap-fish/clay/pkg/bundle"
	"github.com/leap-fish/clay/pkg/components/audio"
	"github.com/leap-fish/clay/pkg/components/spatial"
	"github.com/leap-fish/clay/pkg/components/sprite"
	"github.com/leap-fish/clay/pkg/plugins"
	m "github.com/yohamta/donburi/features/math"
	"math"
	"time"
)

//go:embed assets
var EditorFiles embed.FS

var imageSprite = bundle.New().
	With(spatial.TransformComponent, spatial.Transform{}).
	With(sprite.Component, sprite.Sprite{
		Path: "image:image",
	})

var audioEffect = bundle.New().
	With(audio.Component, audio.SoundEffect{
		Path:   "sfx:bell",
		Volume: 0.02,
	})

type ExamplePlugin struct{}

func (e *ExamplePlugin) Ready(core *pkg.Core) {
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

	go func() {
		time.Sleep(1 * time.Second)
		audioEffect.Spawn(core.World)

		time.Sleep(2 * time.Second)
		audioEffect.Spawn(core.World)

		time.Sleep(2500 * time.Millisecond)
		audioEffect.Spawn(core.World)
	}()
}

func (e *ExamplePlugin) Build(core *pkg.Core) {
	core.Plugin(
		plugins.DefaultPlugins(EditorFiles)...,
	)
	core.SubSystem(&ExampleSystem{})
}
