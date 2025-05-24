package example

import (
	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	s arcadegame.Scene
}

func (g *Game) Init() {
}

func (g *Game) Update() error {
	g.s.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.s.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *Game) Enter() {
	s := &SceneMain{}
	s.Init(g)
	g.SetScene(s)
}

func (g *Game) SetScene(ns arcadegame.Scene) {
	if g.s != nil {
		g.s.Exit()
	}
	g.s = ns
	g.s.Enter()
}

func (g *Game) Exit() {

}
