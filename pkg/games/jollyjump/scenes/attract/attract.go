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
	bgColor color.Color
}

func (s *SceneAttract) Init(r arcadegame.GameRoot) {
	s.r = r
	s.bgColor = color.RGBA{255, 0, 0, 255}
}

func (s *SceneAttract) Enter() {
}

func (s *SceneAttract) Update() error {
	if utils.IsAnyKeyPressed() {
		s.Exit()
		s.Root().SetScene(scenes.SceneTitleToken)
	}
	return nil
}

func (s *SceneAttract) Draw(sc *ebiten.Image) {
	sc.Fill(s.bgColor)
}

func (s *SceneAttract) Exit() {
}

func (s *SceneAttract) Root() arcadegame.GameRoot {
	return s.r
}
