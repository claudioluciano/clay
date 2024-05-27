package text

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
	"image/color"
)

var Component = donburi.NewComponentType[Text]()

type Text struct {
	Content bytes.Buffer

	Source *text.GoTextFaceSource
	Size   float64

	// LineHeight is the relative multiplier per line for font size.
	// Set this to 1.0 if you're unsure.
	LineHeight float64

	PrimaryAlign   text.Align
	SecondaryAlign text.Align

	Color color.NRGBA
}
