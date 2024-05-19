package spatial

import (
	"github.com/yohamta/donburi"
	m "github.com/yohamta/donburi/features/math"
	"math"
)

const (
	radianRotationOffsetHalfCircle = math.Pi / 2
)

var TransformComponent = donburi.NewComponentType[Transform](DefaultTransform)

var DefaultTransform = Transform{
	Position: m.Vec2{},
	Index:    0,
	Rotation: 0,
	Scale:    1.0,
}

type Transform struct {
	Position m.Vec2
	Index    int

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
