package game

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/leap-fish/clay"
	"github.com/leap-fish/clay/config"
	"github.com/yohamta/donburi"
)

type GameAppProvider struct {
	w donburi.World

	options *config.LaunchOptions

	game *clay.ClayGame
}

func NewGameAppProvider(options config.LaunchOptions) *GameAppProvider {
	return &GameAppProvider{
		options: &options,
	}
}

func (g GameAppProvider) Run(world donburi.World, subSystems *clay.SubSystemRegistry, plugins *clay.PluginRegistry) {
	// Initializes the game instance
	g.game = clay.NewClayGame(world, subSystems, plugins, g.options)
	g.game.Init()

	err := ebiten.RunGame(g.game)
	if err != nil {
		panic(err)
	}
}
