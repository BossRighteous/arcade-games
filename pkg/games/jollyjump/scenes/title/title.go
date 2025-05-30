package title

import (
	"image/color"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/scenes"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/settings"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type SceneTitle struct {
	r       arcadegame.GameRoot
	bgColor color.Color
	idleF   int
}

func (s *SceneTitle) Init(r arcadegame.GameRoot) {
	s.r = r
	s.bgColor = color.RGBA{97, 162, 255, 255}
	s.idleF = 0
}

func (s *SceneTitle) Enter() {

}

func (s *SceneTitle) Update() error {
	s.idleF++
	if s.idleF > 1200 {
		s.Exit()
		s.Root().SetScene(scenes.SceneAttractToken)
	}

	if inpututil.IsKeyJustPressed(settings.Esc) {
		//s.Root().SetScene()
		s.Exit()
		return ebiten.Termination
	}
	return nil
}

func (s *SceneTitle) Draw(sc *ebiten.Image) {
	sc.Fill(s.bgColor)
}

func (s *SceneTitle) Exit() {

}

func (s *SceneTitle) Root() arcadegame.GameRoot {
	return s.r
}
