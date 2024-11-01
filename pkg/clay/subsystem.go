package clay

import (
	"cmp"
	"reflect"
	"slices"
	"time"

	log "github.com/leap-fish/clay/pkg/logger"
	"github.com/leap-fish/clay/pkg/render"
	"github.com/yohamta/donburi"
)

type SubSystem interface{}

type Initializable interface {
	Init(w donburi.World)
}

type Renderable interface {
	Render(rg *render.RenderGraph, w donburi.World)
}

type Updatable interface {
	Update(w donburi.World, dt time.Duration)
}

type FullSubSystem interface {
	Initializable
	Updatable
	Renderable
}

type SubSystemRegistry struct {
	SubSystems []SubSystem

	Initializables []Initializable
	Renderables    []Renderable
	Updatables     []Updatable
}

func sortSubSystemSlice[T SubSystem](in []T) []T {
	var systems []T
	systems = append(systems, in...)
	slices.SortFunc(systems, func(a, b T) int {
		var aValue int
		var bValue int

		aOrdered, aOk := SubSystem(a).(Sortable)
		if aOk {
			aValue = aOrdered.Order()
		}

		bOrdered, bOk := SubSystem(b).(Sortable)
		if bOk {
			bValue = bOrdered.Order()
		}

		return cmp.Compare(aValue, bValue)
	})
	return systems
}

func (sr *SubSystemRegistry) Add(systems []SubSystem) {
	// Adds and sorts each system, according to whichever interfaces they actually implement.
	for _, system := range systems {
		initializable, canInit := system.(Initializable)
		if canInit {
			sr.Initializables = append(sr.Initializables, initializable)
			sr.Initializables = sortSubSystemSlice(sr.Initializables)
		}

		renderable, canRender := system.(Renderable)
		if canRender {
			sr.Renderables = append(sr.Renderables, renderable)
			sr.Renderables = sortSubSystemSlice(sr.Renderables)

		}

		updatable, canUpdate := system.(Updatable)
		if canUpdate {
			sr.Updatables = append(sr.Updatables, updatable)
			sr.Updatables = sortSubSystemSlice(sr.Updatables)
		}

		sr.SubSystems = append(sr.SubSystems, system)
		sr.SubSystems = sortSubSystemSlice(sr.SubSystems)

		log.Trace().
			Caller().
			Field(
				log.Field("canInit", canInit),
				log.Field("canRender", canRender),
				log.Field("canUpdate", canUpdate),
			).
			Msgf("Registered subsystem %s", reflect.TypeOf(system))
	}
}
