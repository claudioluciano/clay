package render

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/components/camera"
	"github.com/leap-fish/clay/components/sprite"
	"github.com/leap-fish/clay/components/transform"
	"github.com/leap-fish/clay/render"
	"github.com/leap-fish/clay/resource"
	"github.com/yohamta/donburi"
	"github.com/yohamta/donburi/ecs"
	"github.com/yohamta/donburi/filter"
	"image"
	"time"
)

type DefaultRendererModule struct {
	ecs *ecs.ECS

	cam         *camera.Camera
	spriteQuery *donburi.Query
}

func (d *DefaultRendererModule) Build(core *clay.Core) {
	d.ecs = core.ECS
	d.cam = camera.GetCamera(d.ecs.World)
	d.spriteQuery = donburi.NewQuery(filter.Contains(sprite.Component, transform.Component))
}

func (d *DefaultRendererModule) Ready(core *clay.Core) {

}

func (d *DefaultRendererModule) Update(dt time.Duration) {
}

func (d *DefaultRendererModule) Draw(screen *ebiten.Image) {
	d.spriteQuery.Each(d.ecs.World, func(entry *donburi.Entry) {
		spr := sprite.Component.Get(entry)
		tf := transform.Component.Get(entry)
		if spr.Source == nil {
			res := resource.Get[image.Image](spr.Path)
			spr.Source = ebiten.NewImageFromImage(res)
		}

		render.Draw(spr.Source, render.RenderModeCanvas, tf.Index).
			Scale(tf.Scale).
			OriginMul(spr.Origin.XY()).
			Rotation(tf.Rotation).
			Position(tf.Position.XY()).
			Queue()
	})
}
