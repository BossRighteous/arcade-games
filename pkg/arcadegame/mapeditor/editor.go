package mapeditor

import (
	"fmt"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
}

func (g *Game) Init() {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DebugPrint(
		screen,
		fmt.Sprintf("TPS: %.2f, FPS: %.2f", ebiten.ActualTPS(), ebiten.ActualFPS()),
	)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ScreenWidth, ScreenWidth
}

func (g *Game) Enter() {
}

func (g *Game) SetScene(t arcadegame.SceneToken) {
}

func (g *Game) Exit() {
}
