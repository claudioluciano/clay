package audio

import (
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/leap-fish/clay/resource"
	"github.com/yohamta/donburi"
)

var Component = donburi.NewComponentType[SoundEffect]()

// SoundEffect is an audio player which removes itself once it has finished playing.
type SoundEffect struct {
	Path   resource.Path
	Volume float64

	Player *audio.Player
}
