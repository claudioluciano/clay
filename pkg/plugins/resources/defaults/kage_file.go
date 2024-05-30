package defaults

import (
	"github.com/hajimehoshi/ebiten/v2"
	"io"
)

type KageDefaultHandler struct{}

func (h *KageDefaultHandler) Load(reader io.ReadCloser) (any, error) {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	shader, err := ebiten.NewShader(bytes)
	if err != nil {
		return nil, err
	}

	return shader, nil
}
