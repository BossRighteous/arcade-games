package arcadegame

import "github.com/hajimehoshi/ebiten/v2"

type GameRoot interface {
	ebiten.Game
	Enter()
	SetScene(SceneToken)
	Exit()
}
