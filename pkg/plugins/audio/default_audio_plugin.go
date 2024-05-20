package audio

import (
	"github.com/leap-fish/clay/pkg"
)

type DefaultAudioPlugin struct{}

func (d *DefaultAudioPlugin) Build(core *pkg.Core) {
	core.SubSystem(NewDefaultAudioEffectSystem())
}

func (d *DefaultAudioPlugin) Ready(core *pkg.Core) {
}
