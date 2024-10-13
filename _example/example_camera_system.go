package _example

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/pkg/clay"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/util/ecsutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/math"
	"time"
)

type ExampleCameraSystem struct{}

func (e *ExampleCameraSystem) Build(core *clay.Core) {
}

func (e *ExampleCameraSystem) Ready(core *clay.Core) {
}

func (e *ExampleCameraSystem) Update(w donburi.World, dt time.Duration) {
	cam := ecsutil.FirstOf[camera.Camera](camera.Component, w)
	if cam == nil {
		return
	}

	delta := float64(dt.Milliseconds())
	speed := 1 * delta

	pos := cam.Position

	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		offset := math.NewVec2(0, -speed)
		pos = pos.Add(offset)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		offset := math.NewVec2(0, speed)
		pos = pos.Add(offset)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		offset := math.NewVec2(-speed, 0)
		pos = pos.Add(offset)
	}

	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		offset := math.NewVec2(speed, 0)
		pos = pos.Add(offset)
	}
	cam.SetPosition(pos.XY())
}
