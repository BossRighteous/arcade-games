package nodes

import (
	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/uniqueid"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	TileSize int = 16
)

func T(n int) int {
	return n * TileSize
}

type Node interface {
	ID() uniqueid.UID
	Init()
	Enter()
	Update()
	Draw(*ebiten.Image)
	Exit()
	Delete()
	Scene() arcadegame.Scene
	Root() arcadegame.GameRoot
	Hooks() *LifecycleHooks
	Props() any //Should be pointer to struct
}
