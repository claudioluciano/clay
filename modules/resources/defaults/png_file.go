package defaults

import (
	"image/png"
	"io"
)

type PngDefaultHandler struct{}

func (p *PngDefaultHandler) Load(reader io.ReadCloser) (any, error) {
	pngFile, err := png.Decode(reader)
	return pngFile, err
}
