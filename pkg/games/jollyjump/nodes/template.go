package nodes

import (
	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/hajimehoshi/ebiten/v2"
)

type NodeTemplate struct {
	XScene arcadegame.Scene
	XID    arcadegame.UID
}

func (n *NodeTemplate) ID() arcadegame.UID {
	return n.XID
}

func (n *NodeTemplate) Init() {

}

func (n *NodeTemplate) Enter() {

}

func (n *NodeTemplate) Update() {

}

func (n *NodeTemplate) Draw(s *ebiten.Image) {

}

func (n *NodeTemplate) Exit() {

}

func (n *NodeTemplate) Delete() {

}

func (n *NodeTemplate) Scene() arcadegame.Scene {
	return n.XScene
}
func (n *NodeTemplate) Root() arcadegame.GameRoot {
	return n.XScene.Root()
}
func (n *NodeTemplate) Props() any {
	return n
}
func (n *NodeTemplate) Hooks() *arcadegame.LifecycleHooks {
	return nil
}
