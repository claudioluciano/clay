package editor

import (
	"bytes"
	"embed"
	"flag"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/bundle"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/sprite"
	"github.com/leap-fish/clay/components/transform"
	"github.com/leap-fish/clay/modules/render"
	"github.com/leap-fish/clay/modules/resources"
	render2 "github.com/leap-fish/clay/render"
	"github.com/leap-fish/clay/resource"
	log "github.com/sirupsen/logrus"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/features/debug"
	"github.com/yohamta/donburi/features/math"
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
	ecs  *ecs.ECS
}

var imageSprite = bundle.New().
	With(transform.Component, transform.Transform{}).
	With(sprite.Component, sprite.Sprite{
		Path: "image:image",
	})

func (e *EditorModule) Update(dt time.Duration) {
}

func (e *EditorModule) Draw(screen *ebiten.Image) {
	render2.QueueFunc(func(screen *ebiten.Image, camera *camera.Camera) {
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

		var out bytes.Buffer
		for _, c := range debug.GetEntityCounts(e.ecs.World) {
			out.WriteString(fmt.Sprintf("> %s", c.String()))
			out.WriteString("\n")
		}
		ebitenutil.DebugPrintAt(screen, out.String(), 0, 0)

	}, 0)

}

func (e *EditorModule) Ready(core *clay.Core) {
	e.font = resource.Get[*text.GoTextFaceSource]("font:BaiJamjuree-Regular")

	ent := imageSprite.Spawn(e.ecs.World)
	imageEntry := e.ecs.World.Entry(ent)
	transform.Component.Set(imageEntry, &transform.Transform{
		Position: math.Vec2{240, 260},
		Scale:    0.3,
	})
}

func (e *EditorModule) Build(core *clay.Core) {
	core.Module(
		resources.NewDefaultResourcesModule("assets", EditorFiles),
		&render.DefaultRendererModule{},
	)

	e.ecs = core.ECS

	core.LaunchOptions(clay.LaunchOptions{
		WindowWidth:  *windowWidthFlag,
		WindowHeight: *windowHeightFlag,
	})
	log.Tracef("Window Size set to %dx%d", *windowWidthFlag, *windowHeightFlag)
}
