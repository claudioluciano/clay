package ecsutil

import "github.com/yohamta/donburi"

func FirstOf[T any](component *donburi.ComponentType[T], w donburi.World) *T {
	ent, exists := component.First(w)
	if !exists || ent == nil {
		return nil
	}

	instance := component.Get(ent)
	if instance == nil {
		return nil
	}

	return instance
}
