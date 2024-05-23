package _example

import (
	"bytes"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay/pkg/bundle"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/spatial"
	"github.com/leap-fish/clay/pkg/components/sprite"
	txt "github.com/leap-fish/clay/pkg/components/text"
	"github.com/leap-fish/clay/pkg/render"
	"github.com/leap-fish/clay/pkg/resource"
	"github.com/leap-fish/clay/pkg/util/ecsutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/debug"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
	"image/color"
	"time"
)

var DebugMarker = donburi.NewTag()

type ExampleSystem struct {
	font *text.GoTextFaceSource
}

var movableSprite = bundle.New().
	With(spatial.TransformComponent, spatial.Transform{Scale: 0.1}).
	With(sprite.Component, sprite.Sprite{
		Path: "image:image",
	})

func (e *ExampleSystem) Update(w donburi.World, dt time.Duration) {
	cam := ecsutil.FirstOf(camera.Component, w)
	x, y := ebiten.CursorPosition()
	wposX, wposY := cam.GetWorldCoords(float64(x), float64(y))
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		s := movableSprite.Spawn(w)
		tf := spatial.TransformComponent.Get(w.Entry(s))
		tf.Position = math.NewVec2(wposX, wposY)
	}

	q := donburi.NewQuery(filter.Contains(txt.Component, spatial.TransformComponent))
	entry, exists := q.First(w)
	if !exists || entry == nil {
		return
	}

	t := txt.Component.Get(entry)
	t.Content.Reset()
	t.Content.WriteString(fmt.Sprintf("%0.1f FPS, %0.1f TPS\n\n%dx%d (world: %0fx%0.0f)\n", ebiten.ActualFPS(), ebiten.ActualTPS(), x, y, wposX, wposY))
	for _, c := range debug.GetEntityCounts(w) {
		t.Content.WriteString(fmt.Sprintf("> %s\n", c.String()))
	}
}

func (e *ExampleSystem) Init(w donburi.World) {
	e.font = resource.Get[*text.GoTextFaceSource]("font:BaiJamjuree-Regular")

	DebugMarker.SetName("DebugMarker")

	textBundle := bundle.New().
		With(spatial.TransformComponent, spatial.DefaultTransform).
		With(txt.Component, txt.Text{
			Source:     e.font,
			Content:    bytes.Buffer{},
			Size:       16,
			LineHeight: 1.0,
			Color:      color.RGBA{255, 255, 255, 255},
		}).
		With(DebugMarker, struct{}{})

	textBundle.Spawn(w)

	secondTextBuf := bytes.Buffer{}
	secondTextBuf.WriteString("Clay Engine in golang")
	secondText := bundle.New().
		With(spatial.TransformComponent, spatial.Transform{Position: math.NewVec2(200, 200), Scale: 1.0}).
		With(txt.Component, txt.Text{
			Source:     e.font,
			Content:    secondTextBuf,
			Size:       16,
			LineHeight: 1.0,
			Color:      color.RGBA{255, 100, 100, 255},
		}).
		With(DebugMarker, struct{}{})

	secondText.Spawn(w)
}

func (e *ExampleSystem) Render(rg *render.RenderGraph, w donburi.World) {}
