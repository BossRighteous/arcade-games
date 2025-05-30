package jollyjump

import (
	"fmt"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/scenes"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/scenes/attract"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/scenes/title"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

/*
	Need to tag available scenes as portable for loading from other scenes
*/

type Game struct {
	s       arcadegame.Scene
	queuedS arcadegame.Scene
}

func (g *Game) Init() {
}

func (g *Game) Update() error {
	if g.queuedS != nil {
		g.s = g.queuedS
		g.queuedS = nil
		g.s.Enter()
	}
	err := g.s.Update()
	if err != nil {
		g.Exit()
	}
	return err
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.s.Draw(screen)

	if settings.Debug {
		ebitenutil.DebugPrint(
			screen,
			fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
		)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) Enter() {
	s := &title.SceneTitle{}
	s.Init(g)
	g.s = s
	s.Enter()
}

func (g *Game) SetScene(t arcadegame.SceneToken) {
	if g.s != nil {
		g.s.Exit()
	}
	g.queuedS = provideSceneFromToken(t)
}

func (g *Game) Exit() {
	g.s = nil
}

func provideSceneFromToken(t arcadegame.SceneToken) arcadegame.Scene {
	switch t {
	case scenes.SceneTitleToken:
		return &title.SceneTitle{}
	case scenes.SceneAttractToken:
		return &attract.SceneAttract{}
	default:
		return &title.SceneTitle{}
	}
}
