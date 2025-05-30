package arcadegame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneToken int

type Scene interface {
	Init(GameRoot)
	Enter()
	Update() error
	Draw(screen *ebiten.Image)
	Exit()
	Root() GameRoot
}
