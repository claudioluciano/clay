package text

import (
	"bytes"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/yohamta/donburi"
	"image/color"
)

var Component = donburi.NewComponentType[Text]()

type Text struct {
	FontFace *text.GoTextFace
	Content  bytes.Buffer

	Size int

	// LineHeight is the relative multiplier per line for font size.
	// Set this to 1.0 if you're unsure.
	LineHeight float64

	Color color.RGBA
}
