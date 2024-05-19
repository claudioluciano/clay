package dpi

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/util/ecsutil"
	"github.com/yohamta/donburi"
)

// Avoid using the world lookup by using this local cached one
// Since we're only going to have one of these components
var cachedScaleFactor *float64

type DpiScaleFactor struct {
	ScaleFactor float64
}

var Component = donburi.NewComponentType[DpiScaleFactor](
	DpiScaleFactor{ebiten.Monitor().DeviceScaleFactor()},
)

func GetScaleFactor(w donburi.World) float64 {
	if cachedScaleFactor != nil {
		return *cachedScaleFactor
	}

	scaler := ecsutil.FirstOf(Component, w)
	if scaler == nil {
		return 1.0
	}

	cachedScaleFactor = &scaler.ScaleFactor
	return *cachedScaleFactor
}
