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
	g.s = s
	s.Enter()
}

func (g *Game) SetScene(st arcadegame.SceneToken) {
}

func (g *Game) Exit() {

}
