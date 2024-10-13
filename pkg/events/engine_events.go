package events

import "github.com/yohamta/donburi/features/events"

type WindowSizeUpdate struct {
	Width  int
	Height int
}

var EngineWindowSizeUpdated = events.NewEventType[WindowSizeUpdate]()

var ResourcePluginLoaded = events.NewEventType[int]()
