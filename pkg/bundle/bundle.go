package bundle

import (
	"fmt"
	"github.com/yohamta/donburi"
	"reflect"
	"unsafe"
)

type Bundle struct {
	componentsList []donburi.IComponentType
	components     map[donburi.IComponentType]any
}

func (a *Bundle) With(comp donburi.IComponentType, instance any) *Bundle {
	a.components[comp] = instance
	a.componentsList = append(a.componentsList, comp)
	return a
}

func (a *Bundle) Spawn(w donburi.World) donburi.Entity {
	ent := w.Create(a.componentsList...)
	entry := w.Entry(ent)
	for c, instance := range a.components {
		entry.SetComponent(c, a.componentFromVal(c, instance))
	}
	return ent
}

func (a *Bundle) componentFromVal(ctype donburi.IComponentType, value interface{}) unsafe.Pointer {
	if reflect.TypeOf(value) != ctype.Typ() {
		panic(fmt.Sprintf("Type assertion failed %s vs: %s", ctype.Typ().String(), reflect.TypeOf(value)))
	}
	newVal := reflect.New(ctype.Typ()).Elem()
	newVal.Set(reflect.ValueOf(value))
	ptr := unsafe.Pointer(newVal.UnsafeAddr())

	return ptr
}

func New() *Bundle {
	return &Bundle{
		components: make(map[donburi.IComponentType]any),
	}
}
