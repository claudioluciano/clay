package clay

import (
	"cmp"
	"flag"
	"github.com/hajimehoshi/ebiten/v2"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"reflect"
	"slices"
)

var levelFlag = flag.Int("logging", int(log.InfoLevel), "Sets the logging level of the engine in Logrus levels.")
var loggingColors = flag.Bool("logcolors", false, "Whether logging will have colors enabled")

type LaunchOptions struct {
	WindowWidth  int
	WindowHeight int
	RenderScale  int
}

type Core struct {
	ECS *ecs.ECS

	moduleTypes   []reflect.Type
	builtModules  []reflect.Type
	Modules       []Module
	SortedModules []Module

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
	for _, m := range module {
		moduleType := reflect.TypeOf(m)
		if slices.Contains(c.moduleTypes, moduleType) {
			continue
		}

		log.Tracef("Add module %s", moduleType)
		c.moduleTypes = append(c.moduleTypes, moduleType)
		c.Modules = append(c.Modules, m)
		break
	}

	c.SortedModules = c.sortModules()

	c.Build()

	return c
}

// LaunchOptions is used to configure the `LaunchOptions` structure used to define certain defaults.
func (c *Core) LaunchOptions(options LaunchOptions) *Core {
	c.options = &options
	return c
}

// SortedModules returns modules in sorted order according to their Order() function.
func (c *Core) sortModules() []Module {
	var modules []Module
	modules = append(modules, c.Modules...)
	slices.SortFunc(modules, func(a, b Module) int {
		var aValue int
		var bValue int

		aOrdered, aOk := a.(Sortable)
		if aOk {
			aValue = aOrdered.Order()
		}

		bOrdered, bOk := b.(Sortable)
		if bOk {
			bValue = bOrdered.Order()
		}

		return cmp.Compare(aValue, bValue)
	})
	return modules
}

func (c *Core) Build() {
	// Modules are the first things to get initialized
	for _, module := range c.SortedModules {
		moduleType := reflect.TypeOf(module)
		if slices.Contains(c.builtModules, moduleType) {
			continue
		}
		c.builtModules = append(c.builtModules, moduleType)
		module.Build(c)
	}
}

func (c *Core) Run() {
	log.Trace("Clay: Run()")

	// After all modules are built, run the `Ready` function
	for _, module := range c.SortedModules {
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
