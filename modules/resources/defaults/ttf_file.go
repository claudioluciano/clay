package defaults

import (
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"io"
)

type TtfDefaultHandler struct{}

func (t *TtfDefaultHandler) Load(reader io.ReadCloser) (any, error) {
	font, err := text.NewGoTextFaceSource(reader)
	return font, err
}
