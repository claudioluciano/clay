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

// LaunchOptions is a simple struct that holds a standard set of launch Options that the user may change.
type LaunchOptions struct {

	// Physical window width in pixels.
	WindowWidth int

	// Physical window height in pixels.
	WindowHeight int

	// RenderScale determines the factor at which rendering is scaled at. 2.0 will result in every pixel taking two pixels
	// on screen.
	RenderScale float64

	// UseDPIScaling enables automatic scaling depending on monitor DPI, and is calculated using the ebitengine
	// API. This should make rendering look identical on different type of screens, such as a 4k "Retina" screen
	// compared to a normal 1080p screen. If you are making a typical "pixel game", you may want to leave this `false`,
	//as this can cause rendering to be blurred.
	UseDPIScaling bool

	// VsyncMode will restrict rendering to your monitor refresh rate,
	// or whether to use ebiten's own rendering scheduling.
	VsyncMode bool
}

// DefaultLaunchOptions is just some reasonably sane default launch options, available for use.
var DefaultLaunchOptions = LaunchOptions{
	WindowWidth:   1920,
	WindowHeight:  1080,
	RenderScale:   1,
	UseDPIScaling: true,
	VsyncMode:     true,
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

	Options *LaunchOptions
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
		Options: &LaunchOptions{
			WindowWidth:   800,
			WindowHeight:  600,
			UseDPIScaling: true,
			RenderScale:   1.0,
		},
	}

	core.PluginRegistry = NewPluginRegistry(core)

	return core
}

// Plugin registers a feature providing plugin with the engine.
// This will make sure that the lifecycle code is ran in the correct order.
func (c *Core) Plugin(plugins ...Plugin) *Core {
	c.PluginRegistry.Add(plugins)
	return c
}

// SubSystem adds a system to the subsystem registry, allowing plugins to add features.
// This is normally called from inside a plugin instance that has been added with `Core.Plugin()`, but
// can also be called standalone, from something like a main function.
func (c *Core) SubSystem(systems ...SubSystem) *Core {
	c.SubSystemRegistry.Add(systems)
	return c
}

// LaunchOptions is used to configure the `LaunchOptions` structure
// which is used to define engine defaults.
// In LaunchOptions, are things like window size, render scale and other fundamental options.
func (c *Core) LaunchOptions(options LaunchOptions) *Core {
	c.Options = &options
	return c
}

// build makes sure all plugins get initialized (their build functions are called once, recursively).
// This must be called before `Core.Run()`.
func (c *Core) build() {
	c.PluginRegistry.BuildPlugins()
}

// Run is used to actually launch the underlying Ebitengine instance, with the added clay subsystems and configuration.
// Make sure you call `Core.Build()` beforehand.
func (c *Core) Run() {
	c.build()

	for _, plugin := range c.PluginRegistry.Plugins {
		plugin.Ready(c)
	}

	// Defaults set by ebiten
	ebiten.SetWindowSize(c.Options.WindowWidth, c.Options.WindowHeight)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetVsyncEnabled(c.Options.VsyncMode)

	// Initializes the game instance
	c.Game = &ClayGame{
		Core:  c,
		World: c.World,
	}

	c.Game.Init()
	err := ebiten.RunGame(c.Game)
	if err != nil {
		panic(err)
	}
}
