package nodes

import (
	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/uniqueid"
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

type NodeBase struct {
	XScene arcadegame.Scene
	XID    uniqueid.UID
	XHooks LifecycleHooks
	XProps any
}

func (n *NodeBase) ID() uniqueid.UID {
	return n.XID
}
func (n *NodeBase) Init() {
	if n.XHooks.InitFunc != nil {
		(*n.XHooks.InitFunc)(n)
	}
}
func (n *NodeBase) Enter() {
	if n.XHooks.EnterFunc != nil {
		(*n.XHooks.EnterFunc)(n)
	}
}
func (n *NodeBase) Update() {
	if n.XHooks.UpdateFunc != nil {
		(*n.XHooks.UpdateFunc)(n)
	}
}
func (n *NodeBase) Draw(s *ebiten.Image) {
	if n.XHooks.DrawFunc != nil {
		(*n.XHooks.DrawFunc)(n, s)
	}
}
func (n *NodeBase) Exit() {
	if n.XHooks.ExitFunc != nil {
		(*n.XHooks.ExitFunc)(n)
	}
}
func (n *NodeBase) Delete() {
	if n.XHooks.DeleteFunc != nil {
		(*n.XHooks.DeleteFunc)(n)
	}
}
func (n *NodeBase) Scene() arcadegame.Scene {
	return n.XScene
}
func (n *NodeBase) Root() arcadegame.GameRoot {
	return n.XScene.Root()
}
func (n *NodeBase) Props() any {
	return n.XProps
}
func (n *NodeBase) Hooks() *LifecycleHooks {
	return &n.XHooks
}

func MakeNodeBase(
	scene arcadegame.Scene,
	hooks LifecycleHooks,
	props any,
) *NodeBase {
	id := uniqueid.UniqueID()
	n := &NodeBase{
		XScene: scene,
		XID:    id,
		XHooks: hooks,
		XProps: props,
	}
	return n
}
