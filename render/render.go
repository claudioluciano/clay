package render

// Mode represents the type of rendering that will be done.
// ModeCanvas is used for screen space rendering.
// ModeWorld is used for world space rendering.
type Mode uint

const (
	ModeCanvas Mode = iota + 1
	ModeWorld
)
