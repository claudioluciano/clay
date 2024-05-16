package transform

import (
	"github.com/yohamta/donburi"
	m "github.com/yohamta/donburi/features/math"
	"math"
)

const (
	radianRotationOffsetHalfCircle = math.Pi / 2
)

var Component = donburi.NewComponentType[Transform](Transform{
	Position: m.NewVec2(0, 0),
	Rotation: 0,
	Scale:    1.0,
})

type Transform struct {
	Position m.Vec2
	Rotation float64
	Scale    float64
}

func (w *Transform) Forward() m.Vec2 {
	return m.NewVec2(1, 0).
		Rotate(w.Rotation - radianRotationOffsetHalfCircle)
}

func (w *Transform) Right() m.Vec2 {
	return m.NewVec2(0, 1).
		Rotate(w.Rotation - radianRotationOffsetHalfCircle)
}
