package defaults

import (
	"github.com/hajimehoshi/ebiten/v2/audio/vorbis"
	"github.com/leap-fish/clay/plugins/audio"
	"io"
)

type OggDefaultHandler struct{}

func (o *OggDefaultHandler) Load(reader io.ReadCloser) (any, error) {
	var err error
	s, err := vorbis.DecodeWithoutResampling(reader)
	if err != nil {
		return nil, err
	}
	bytes, err := io.ReadAll(s)
	if err != nil {
		return nil, err
	}
	return audio.SoundBytes(bytes), nil
}
