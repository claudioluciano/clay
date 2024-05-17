package clay

import (
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var levelFlag = flag.Int("logging", int(log.InfoLevel), "Sets the logging level of the engine in Logrus levels (0 to 6).")
var loggingColors = flag.Bool("logcolors", false, "Whether logging will have colors enabled")

// LaunchOptions is a simple struct that holds a standard set of launch options that the user may change.
type LaunchOptions struct {
	WindowWidth  int
	WindowHeight int
	RenderScale  int
}

// Core holds subsystems for Clay.
// This struct contains most relevant methods of implementing features for a Clay based application.
type Core struct {
	// The root ECS instance
	// TODO: Does this actually need to be exposed?
	ECS *ecs.ECS

	// Shortcut to ECS.World
	World donburi.World

	// Used for rendering
	RenderGraph *RenderGraph

	PluginRegistry    *PluginRegistry
	SubSystemRegistry *SubSystemRegistry

	Game *ClayGame

	options *LaunchOptions
}

func New() *Core {
	flag.Parse()

	log.SetLevel(log.Level(*levelFlag))
	log.SetFormatter(&log.TextFormatter{ForceColors: *loggingColors})

	world := donburi.NewWorld()
	ecsInstance := ecs.NewECS(world)

	core := &Core{
		ECS:               ecsInstance,
		World:             world,
		RenderGraph:       &RenderGraph{},
		SubSystemRegistry: &SubSystemRegistry{},
		options: &LaunchOptions{
			WindowWidth:  800,
			WindowHeight: 600,
			RenderScale:  1.0,
		},
	}

	core.PluginRegistry = NewPluginRegistry(core)

	return core
}

func (c *Core) Plugin(plugins ...Plugin) *Core {
	c.PluginRegistry.Add(plugins)
	return c
}

func (c *Core) SubSystem(systems ...SubSystem) *Core {
	c.SubSystemRegistry.Add(systems)
	return c
}

// LaunchOptions is used to configure the `LaunchOptions` structure used to define certain defaults.
func (c *Core) LaunchOptions(options LaunchOptions) *Core {
	c.options = &options
	return c
}

func (c *Core) Build() {
	c.PluginRegistry.BuildPlugins()
}

func (c *Core) Run() {
	for _, plugin := range c.PluginRegistry.Plugins {
		plugin.Ready(c)
	}

	// Defaults
	ebiten.SetWindowSize(c.options.WindowWidth, c.options.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	// Initializes the game instance
	c.Game = &ClayGame{
		RenderScale: 1.0,
		Core:        c,
		World:       c.World,
	}

	err := ebiten.RunGame(c.Game)
	if err != nil {
		panic(err)
	}
}
