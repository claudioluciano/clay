package main

import (
	"flag"

	"_example"

	"github.com/leap-fish/clay/pkg/clay"
	"github.com/leap-fish/clay/pkg/config"
	"github.com/leap-fish/clay/pkg/game"
	log "github.com/leap-fish/clay/pkg/logger"
)

var (
	windowWidthFlag  = flag.Int("width", 800, "window width")
	windowHeightFlag = flag.Int("height", 600, "window height")
)

func main() {
	// log.SetLogger(_example.NewMyLogger())

	c := clay.New()
	c.Provider(game.NewGameAppProvider(config.LaunchOptions{
		WindowWidth:   *windowWidthFlag,
		WindowHeight:  *windowHeightFlag,
		UseDPIScaling: true,
		RenderScale:   1,
		VsyncMode:     true,
	}))
	log.Trace().Caller().Msgf("Window Size set to %dx%d", *windowWidthFlag, *windowHeightFlag)
	c.Plugin(&_example.ExamplePlugin{})
	c.Run()
}
