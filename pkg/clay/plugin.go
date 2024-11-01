package clay

import (
	"cmp"
	"reflect"
	"slices"

	log "github.com/leap-fish/clay/pkg/logger"
)

type PluginRegistry struct {
	Plugins []Plugin

	pluginTypes  []reflect.Type
	builtPlugins []reflect.Type

	Core *Core
}

func NewPluginRegistry(core *Core) *PluginRegistry {
	return &PluginRegistry{Core: core}
}

func (pr *PluginRegistry) Add(plugin []Plugin) {
	for _, m := range plugin {
		pluginType := reflect.TypeOf(m)
		if slices.Contains(pr.pluginTypes, pluginType) {
			continue
		}

		log.Trace().
			Caller().
			Msgf("Add plugin %s", pluginType)
		pr.pluginTypes = append(pr.pluginTypes, pluginType)
		pr.Plugins = append(pr.Plugins, m)
	}

	pr.Plugins = pr.sortPlugins()

	pr.BuildPlugins()
}

func (pr *PluginRegistry) BuildPlugins() {
	for _, plugin := range pr.Plugins {
		pluginType := reflect.TypeOf(plugin)
		if slices.Contains(pr.builtPlugins, pluginType) {
			continue
		}
		pr.builtPlugins = append(pr.builtPlugins, pluginType)
		plugin.Build(pr.Core)
	}
}

func (pr *PluginRegistry) sortPlugins() []Plugin {
	var plugins []Plugin
	plugins = append(plugins, pr.Plugins...)
	slices.SortFunc(plugins, func(a, b Plugin) int {
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
	return plugins
}

// Plugin is the interface that allows the Clay engine to manage subsystems.
// Implement `Plugin` for your custom module, and add it using `Core.Plugin()`.
type Plugin interface {
	Build(core *Core)
	Ready(core *Core)
}
