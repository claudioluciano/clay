package main

import (
	"_example"
	"flag"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/config"
	"github.com/leap-fish/clay/game"
	log "github.com/sirupsen/logrus"
)

var (
	windowWidthFlag  = flag.Int("width", 800, "window width")
	windowHeightFlag = flag.Int("height", 600, "window height")
)

func main() {
	c := clay.New()
	c.Provider(game.NewGameAppProvider(config.LaunchOptions{
		WindowWidth:   *windowWidthFlag,
		WindowHeight:  *windowHeightFlag,
		UseDPIScaling: true,
		RenderScale:   1.0,
		VsyncMode:     true,
	}))
	log.Tracef("Window Size set to %dx%d", *windowWidthFlag, *windowHeightFlag)
	c.Plugin(&_example.ExamplePlugin{})
	c.Run()
}
