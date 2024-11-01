package _example

import (
	"bytes"
	"fmt"
	"image/color"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/leap-fish/clay/pkg/bundle"
	"github.com/leap-fish/clay/pkg/components/animsprite"
	"github.com/leap-fish/clay/pkg/components/camera"
	"github.com/leap-fish/clay/pkg/components/spatial"
	"github.com/leap-fish/clay/pkg/components/sprite"
	txt "github.com/leap-fish/clay/pkg/components/text"
	"github.com/leap-fish/clay/pkg/events"
	log "github.com/leap-fish/clay/pkg/logger"
	"github.com/leap-fish/clay/pkg/render"
	"github.com/leap-fish/clay/pkg/util/ecsutil"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/features/debug"
	"github.com/yohamta/donburi/features/math"
	"github.com/yohamta/donburi/filter"
)

var DebugMarker = donburi.NewTag("DebugMarker")

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
	worldPos := math.NewVec2(wposX, wposY)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButton0) {
		s := movableSprite.Spawn(w)
		tf := spatial.TransformComponent.Get(w.Entry(s))
		tf.Position = worldPos
	}

	q := donburi.NewQuery(filter.Contains(txt.Component, spatial.TransformComponent))
	entry, exists := q.First(w)
	if !exists || entry == nil {
		return
	}

	t := txt.Component.Get(entry)
	t.Content.Reset()
	t.Content.WriteString(
		fmt.Sprintf(
			"%0.1f FPS, %0.1f TPS\n\n%dx%d (world: %#v)\n\nCamera pos: %#v\n\n",
			ebiten.ActualFPS(),
			ebiten.ActualTPS(),
			x,
			y,
			worldPos,
			cam.Position,
		),
	)
	for _, c := range debug.GetEntityCounts(w) {
		t.Content.WriteString(fmt.Sprintf("> %s\n", c.String()))
	}
}

func (e *ExampleSystem) Init(w donburi.World) {
	events.ResourcePluginLoaded.Subscribe(w, func(w donburi.World, event int) {
		log.Info().Msg("Spawning spritesheet")
		spriteSheet := bundle.New().
			With(spatial.TransformComponent, spatial.Transform{Scale: 5.0, Position: math.NewVec2(20, 20)}).
			With(animsprite.Component, animsprite.New(
				"image:spritesheet",
				"idle",
				animsprite.SpriteSheetSize{
					FrameHeight: 32,
					FrameWidth:  32,
					ImageWidth:  416,
					ImageHeight: 256,
				},
				map[string]*animsprite.Animation{
					"idle": animsprite.NewAnimation(time.Millisecond*120, "1-13", 1),
					"run":  animsprite.NewAnimation(time.Millisecond*50, "1-8", 2),
				},
				ebiten.FilterNearest,
			))

		spriteSheet.Spawn(w)

		go func() {
			spr := ecsutil.FirstOf[animsprite.AnimSprite](animsprite.Component, w)
			time.Sleep(3 * time.Second)
			spr.CurrentAnimation = "run"

			time.Sleep(10 * time.Second)
			spr.CurrentAnimation = "idle"
		}()
	})

	DebugMarker.SetName("DebugMarker")

	textBundle := bundle.New().
		With(spatial.TransformComponent, spatial.DefaultTransform).
		With(txt.Component, txt.Text{
			Path:       "font:BaiJamjuree-Regular",
			Content:    bytes.Buffer{},
			Size:       16,
			LineHeight: 1.0,
			Color:      color.NRGBA{255, 255, 255, 255},
		}).
		With(DebugMarker, donburi.Tag("DebugMarker"))

	textBundle.Spawn(w)

	secondTextBuf := bytes.Buffer{}
	secondTextBuf.WriteString("Clay Engine in golang")
	secondText := bundle.New().
		With(spatial.TransformComponent, spatial.Transform{Position: math.NewVec2(200, 200), Scale: 1.0}).
		With(txt.Component, txt.Text{
			Path:       "font:BaiJamjuree-Regular",
			Content:    secondTextBuf,
			Size:       16,
			LineHeight: 1.0,
			Color:      color.NRGBA{255, 100, 100, 255},
		}).
		With(DebugMarker, donburi.Tag("DebugMarker"))

	secondText.Spawn(w)
}

func (e *ExampleSystem) Render(rg *render.RenderGraph, w donburi.World) {
}
