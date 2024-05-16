package clay

import (
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var levelFlag = flag.Int("logging", int(log.InfoLevel), "Sets the logging level of the engine in Logrus levels.")

type LaunchOptions struct {
	WindowWidth  int
	WindowHeight int
}

type Core struct {
	ECS     *ecs.ECS
	Modules []Module

	options *LaunchOptions
}

func (c *Core) Update() error {
	return nil
}

func (c *Core) Draw(screen *ebiten.Image) {
}

func (c *Core) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}

func New() *Core {
	flag.Parse()

	log.SetLevel(log.Level(*levelFlag))

	return &Core{
		ECS: ecs.NewECS(donburi.NewWorld()),
		options: &LaunchOptions{
			WindowWidth:  800,
			WindowHeight: 600,
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
		module.Init(c)
	}

	// Defaults
	ebiten.SetWindowSize(c.options.WindowWidth, c.options.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	err := ebiten.RunGame(c)
	if err != nil {
		panic(err)
	}
}
