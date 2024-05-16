package editor

import (
	"flag"
	"github.com/leap-fish/clay"
	log "github.com/sirupsen/logrus"
	"os"
)

var (
	windowWidthFlag  = flag.Int("width", 800, "window width")
	windowHeightFlag = flag.Int("height", 600, "window height")
)

type EditorModule struct {
}

func (e *EditorModule) Init(core *clay.Core) {
	args := os.Args
	log.
		WithField("Args", args).
		Debug("Clay Editor starting")

	core.LaunchOptions(clay.LaunchOptions{
		WindowWidth:  *windowWidthFlag,
		WindowHeight: *windowHeightFlag,
	})
	log.Tracef("Window Size set to %dx%d", *windowWidthFlag, *windowHeightFlag)
}
