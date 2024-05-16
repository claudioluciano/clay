package clay

import (
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var levelFlag = flag.Int("logging", int(log.InfoLevel), "Sets the logging level of the engine in Logrus levels.")
var loggingColors = flag.Bool("logcolors", false, "Whether logging will have colors enabled")

type LaunchOptions struct {
	WindowWidth  int
	WindowHeight int
	RenderScale  int
}

type Core struct {
	ECS     *ecs.ECS
	Modules []Module

	Game *ClayGame

	options *LaunchOptions
}

func New() *Core {
	flag.Parse()

	log.SetLevel(log.Level(*levelFlag))
	log.SetFormatter(&log.TextFormatter{ForceColors: *loggingColors})

	return &Core{
		ECS: ecs.NewECS(donburi.NewWorld()),
		options: &LaunchOptions{
			WindowWidth:  800,
			WindowHeight: 600,
			RenderScale:  1.0,
		},
	}
}

func (c *Core) Module(module ...Module) *Core {
	c.Modules = append(c.Modules, module...)

	return c
}

func (c *Core) LaunchOptions(options LaunchOptions) *Core {
	c.options = &options
	return c
}

func (c *Core) Run() {
	// Modules are the first things to get initialized
	for _, module := range c.Modules {
		module.Build(c)
	}

	// After all modules are built, run the `Ready` function
	for _, module := range c.Modules {
		module.Ready(c)
	}

	// Defaults
	ebiten.SetWindowSize(c.options.WindowWidth, c.options.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Initializes the game instance
	c.Game = &ClayGame{
		RenderScale: 1.0,
		Core:        c,
	}

	err := ebiten.RunGame(c.Game)
	if err != nil {
		panic(err)
	}
}
