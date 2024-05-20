package pkg

import "github.com/yohamta/donburi"

// Sortable lets Clay systems order systems and plugins.
type Sortable interface {
	Order() int
}

type AppProvider interface {
	Run(world donburi.World, subSystems *SubSystemRegistry, plugins *PluginRegistry)
}
