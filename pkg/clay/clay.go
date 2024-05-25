package clay

import (
	"flag"
	"github.com/leap-fish/clay/pkg/config"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
)

var levelFlag = flag.Int("logging", int(log.InfoLevel), "Sets the logging level of the engine in Logrus levels (0 to 6).")
var loggingColors = flag.Bool("logcolors", true, "Whether logging will have colors enabled")

// Core holds subsystems for Clay.
// This struct contains most relevant methods of implementing features for a Clay based application.
type Core struct {
	// The root ECS instance
	// TODO: Does this actually need to be exposed?
	ECS *ecs.ECS

	// Shortcut to ECS.World
	World donburi.World

	PluginRegistry    *PluginRegistry
	SubSystemRegistry *SubSystemRegistry

	Options config.LaunchOptions

	provider AppProvider
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
		SubSystemRegistry: &SubSystemRegistry{},
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

func (c *Core) Provider(provider AppProvider) *Core {
	c.provider = provider
	return c
}

// build makes sure all plugins get initialized (their build functions are called once, recursively).
// This must be called before `Core.Run()`.
func (c *Core) build() {
	c.PluginRegistry.BuildPlugins()
}

// Run is used to actually launch the underlying app provider instance, with the added clay subsystems and configuration.
// Make sure you call `Core.Build()` beforehand.
func (c *Core) Run() {
	c.build()

	for _, plugin := range c.PluginRegistry.Plugins {
		plugin.Ready(c)
	}

	if c.provider == nil {
		log.Fatal("Clay Core.Run() was called without setting AppProvider using Core.Provider() first")
		return
	}

	c.provider.Run(c.World, c.SubSystemRegistry, c.PluginRegistry)
}
