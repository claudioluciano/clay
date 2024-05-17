package render

// RenderMode represents the type of rendering that will be done.
// RenderModeCanvas is used for screen space rendering.
// RenderModeWorld is used for world space rendering.
type RenderMode uint

const (
	RenderModeCanvas RenderMode = iota + 1
	RenderModeWorld
)
