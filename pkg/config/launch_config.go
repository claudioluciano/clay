package config

// LaunchOptions is a simple struct that holds a standard set of launch options that the user may change.
type LaunchOptions struct {

	// Physical window width in pixels.
	WindowWidth int

	// Physical window height in pixels.
	WindowHeight int

	// RenderScale determines the factor at which rendering is scaled at. 2.0 will result in every pixel taking two pixels
	// on screen.
	RenderScale float64

	// UseDPIScaling enables automatic scaling depending on monitor DPI, and is calculated using the ebitengine
	// API. This should make rendering look identical on different type of screens, such as a 4k "Retina" screen
	// compared to a normal 1080p screen. If you are making a typical "pixel game", you may want to leave this `false`,
	//as this can cause rendering to be blurred.
	UseDPIScaling bool

	// VsyncMode will restrict rendering to your monitor refresh rate,
	// or whether to use ebiten's own rendering scheduling.
	VsyncMode bool
}

// DefaultLaunchOptions is just some reasonably sane default launch options, available for use.
var DefaultLaunchOptions = LaunchOptions{
	WindowWidth:   800,
	WindowHeight:  600,
	RenderScale:   1,
	UseDPIScaling: true,
	VsyncMode:     true,
}
