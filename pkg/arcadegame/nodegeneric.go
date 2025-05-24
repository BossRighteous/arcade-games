package arcadegame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

/*
	NodeGeneric is a free-form Node implementation that relies on Hook callbacks and type asserted Props
	pkg\games\example\bgnode\bg.go shows an implementation
*/

type NodeGeneric struct {
	XScene Scene
	XID    UID
	XHooks LifecycleHooks
	XProps any
}

func (n *NodeGeneric) ID() UID {
	return n.XID
}
func (n *NodeGeneric) Init() {
	if n.XHooks.InitFunc != nil {
		(*n.XHooks.InitFunc)(n)
	}
}
func (n *NodeGeneric) Enter() {
	if n.XHooks.EnterFunc != nil {
		(*n.XHooks.EnterFunc)(n)
	}
}
func (n *NodeGeneric) Update() {
	if n.XHooks.UpdateFunc != nil {
		(*n.XHooks.UpdateFunc)(n)
	}
}
func (n *NodeGeneric) Draw(s *ebiten.Image) {
	if n.XHooks.DrawFunc != nil {
		(*n.XHooks.DrawFunc)(n, s)
	}
}
func (n *NodeGeneric) Exit() {
	if n.XHooks.ExitFunc != nil {
		(*n.XHooks.ExitFunc)(n)
	}
}
func (n *NodeGeneric) Delete() {
	if n.XHooks.DeleteFunc != nil {
		(*n.XHooks.DeleteFunc)(n)
	}
}
func (n *NodeGeneric) Scene() Scene {
	return n.XScene
}
func (n *NodeGeneric) Root() GameRoot {
	return n.XScene.Root()
}
func (n *NodeGeneric) Props() any {
	return n.XProps
}
func (n *NodeGeneric) Hooks() *LifecycleHooks {
	return &n.XHooks
}

func MakeNodeGeneric(
	scene Scene,
	hooks LifecycleHooks,
	props any,
) *NodeGeneric {
	id := UniqueID()
	n := &NodeGeneric{
		XScene: scene,
		XID:    id,
		XHooks: hooks,
		XProps: props,
	}
	return n
}
