package attract

import (
	"image/color"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/scenes"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneAttract struct {
	r       arcadegame.GameRoot
	ns      []arcadegame.Node
	bgColor color.Color
}

func (s *SceneAttract) Init(r arcadegame.GameRoot) {
	s.r = r
	s.ns = make([]arcadegame.Node, 0)
	s.bgColor = color.RGBA{255, 0, 0, 255}
}

func (s *SceneAttract) Enter() {
	for _, n := range s.ns {
		n.Enter()
	}
}

func (s *SceneAttract) Update() error {
	for _, n := range s.ns {
		n.Update()
	}

	if utils.IsAnyKeyPressed() {
		s.Exit()
		s.Root().SetScene(scenes.SceneTitleToken)
	}
	return nil
}

func (s *SceneAttract) Draw(sc *ebiten.Image) {
	sc.Fill(s.bgColor)

	for _, n := range s.ns {
		n.Draw(sc)
	}
}

func (s *SceneAttract) Exit() {
	for _, n := range s.ns {
		n.Exit()
		n.Delete()
	}
}

func (s *SceneAttract) Root() arcadegame.GameRoot {
	return s.r
}
