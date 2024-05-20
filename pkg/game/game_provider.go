package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay/pkg"
	"github.com/leap-fish/clay/pkg/config"
	"github.com/yohamta/donburi"
)

type GameAppProvider struct {
	w donburi.World

	options *config.LaunchOptions

	game *pkg.ClayGame
}

func NewGameAppProvider(options config.LaunchOptions) *GameAppProvider {
	return &GameAppProvider{
		options: &options,
	}
}

func (g GameAppProvider) Run(world donburi.World, subSystems *pkg.SubSystemRegistry, plugins *pkg.PluginRegistry) {
	// Initializes the game instance
	g.game = pkg.NewClayGame(world, subSystems, plugins, g.options)
	g.game.Init()

	err := ebiten.RunGame(g.game)
	if err != nil {
		panic(err)
	}
}
