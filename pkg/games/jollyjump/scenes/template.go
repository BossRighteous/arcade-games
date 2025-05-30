package scenes

import (
	"image/color"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneTemplate struct {
	r       arcadegame.GameRoot
	ns      []arcadegame.Node
	bgColor color.Color
}

func (s *SceneTemplate) Init(r arcadegame.GameRoot) {
	s.r = r
	s.ns = make([]arcadegame.Node, 0)
	s.bgColor = color.RGBA{255, 255, 255, 255}
}

func (s *SceneTemplate) Enter() {
	for _, n := range s.ns {
		n.Enter()
	}
}

func (s *SceneTemplate) Update() {
	for _, n := range s.ns {
		n.Update()
	}
}

func (s *SceneTemplate) Draw(sc *ebiten.Image) {
	sc.Fill(s.bgColor)

	for _, n := range s.ns {
		n.Draw(sc)
	}
}

func (s *SceneTemplate) Exit() {
	for _, n := range s.ns {
		n.Exit()
		n.Delete()
	}
}

func (s *SceneTemplate) Root() arcadegame.GameRoot {
	return s.r
}
