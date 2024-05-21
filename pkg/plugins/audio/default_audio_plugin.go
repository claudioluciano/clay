package audio

import "github.com/leap-fish/clay/pkg/clay"

type DefaultAudioPlugin struct{}

func (d *DefaultAudioPlugin) Build(core *clay.Core) {
	core.SubSystem(NewDefaultAudioEffectSystem())
}

func (d *DefaultAudioPlugin) Ready(core *clay.Core) {
}
