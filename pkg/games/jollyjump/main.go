package jollyjump

import (
	"log"

	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

func Main() {
	ebiten.SetWindowSize(settings.WindowWidth, settings.WindowHeight)
	ebiten.SetWindowTitle(settings.WindowTitle)
	ebiten.SetTPS(settings.TPS)
	ebiten.SetVsyncEnabled(settings.VSyncEnabled)

	g := &Game{}
	g.Init()
	g.Enter()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
