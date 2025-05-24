package arcadegame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type LifecycleHooks struct {
	InitFunc   *func(Node)
	EnterFunc  *func(Node)
	UpdateFunc *func(Node)
	DrawFunc   *func(Node, *ebiten.Image)
	ExitFunc   *func(Node)
	DeleteFunc *func(Node)
}

type Node interface {
	ID() UID
	Init()
	Enter()
	Update()
	Draw(*ebiten.Image)
	Exit()
	Delete()
	Scene() Scene
	Root() GameRoot
	Hooks() *LifecycleHooks // Return nil for concrete nodes
	Props() any             // Should be pointer to known struct or nil
}
