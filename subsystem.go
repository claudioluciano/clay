package clay

import (
	"cmp"
	"github.com/leap-fish/clay/render"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi"
	"reflect"
	"slices"
	"time"
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
		clientInitializable, canInit := system.(Initializable)
		if canInit {
			sr.Initializables = append(sr.Initializables, clientInitializable)
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

		log.
			WithField("canInit", canInit).
			WithField("canRender", canRender).
			WithField("canUpdate", canUpdate).
			Tracef("Registered subsystem %s", reflect.TypeOf(system))
	}
}
