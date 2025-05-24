package arcadegame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init(GameRoot)
	Enter()
	Update()
	Draw(screen *ebiten.Image)
	Exit()
	Root() GameRoot
}
