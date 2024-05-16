package editor

import (
	"embed"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/modules/render"
	"github.com/leap-fish/clay/modules/resources"
	"github.com/leap-fish/clay/resource"
	log "github.com/sirupsen/logrus"
	"image/color"
	"time"
)

var (
	windowWidthFlag  = flag.Int("width", 800, "window width")
	windowHeightFlag = flag.Int("height", 600, "window height")
)

//go:embed assets
var EditorFiles embed.FS

type EditorModule struct {
	font *text.GoTextFaceSource
}

func (e *EditorModule) Update(dt time.Duration) {
}

func (e *EditorModule) Draw(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	op.GeoM.Translate(32, 32)
	op.LineSpacing = 32
	text.Draw(
		screen,
		fmt.Sprintf("Clay Editor"),
		&text.GoTextFace{
			Source: e.font,
			Size:   32,
		},
		op,
	)
}

func (e *EditorModule) Ready(core *clay.Core) {
	e.font = resource.Get[*text.GoTextFaceSource]("font:BaiJamjuree-Regular")
}

func (e *EditorModule) Build(core *clay.Core) {
	core.Module(
		resources.NewDefaultResourcesModule("assets", EditorFiles),
		&render.DefaultRendererModule{},
	)

	core.LaunchOptions(clay.LaunchOptions{
		WindowWidth:  *windowWidthFlag,
		WindowHeight: *windowHeightFlag,
	})
	log.Tracef("Window Size set to %dx%d", *windowWidthFlag, *windowHeightFlag)
}
