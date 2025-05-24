package nodes

import (
	"image"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/uniqueid"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/bools"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/floats"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/imgs"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/ints"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/rects"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/strs"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/vm"
	"github.com/hajimehoshi/ebiten/v2"
)

type NodeBase struct {
	XScene      arcadegame.Scene // Public for initial set across packages :/
	XID         uniqueid.UID
	XInitFunc   *func(Node)
	XEnterFunc  *func(Node)
	XUpdateFunc *func(Node)
	XDrawFunc   *func(Node, *ebiten.Image)
	XExitFunc   *func(Node)
	XDeleteFunc *func(Node)
}

func (n *NodeBase) ID() uniqueid.UID {
	return n.XID
}
func (n *NodeBase) Init() {
	if n.XInitFunc != nil {
		(*n.XInitFunc)(n)
	}
}
func (n *NodeBase) Enter() {
	if n.XEnterFunc != nil {
		(*n.XEnterFunc)(n)
	}
}
func (n *NodeBase) Update() {
	if n.XUpdateFunc != nil {
		(*n.XUpdateFunc)(n)
	}
}
func (n *NodeBase) Draw(s *ebiten.Image) {
	if n.XDrawFunc != nil {
		(*n.XDrawFunc)(n, s)
	}
}
func (n *NodeBase) Exit() {
	if n.XExitFunc != nil {
		(*n.XExitFunc)(n)
	}
}
func (n *NodeBase) Delete() {
	if n.XDeleteFunc != nil {
		(*n.XDeleteFunc)(n)
	}
}
func (n *NodeBase) Scene() arcadegame.Scene {
	return n.XScene
}
func (n *NodeBase) Root() arcadegame.GameRoot {
	return n.XScene.Root()
}
func (n *NodeBase) VM() vm.VM {
	return n.XScene.VM()
}

func (n *NodeBase) Bool(pid bools.BoolVar) (bool, bool) {
	return n.XScene.VM().Bool(vars.VMBool{NID: n.XID, PID: pid})
}
func (n *NodeBase) SetBool(pid bools.BoolVar, val bool) {
	n.XScene.VM().SetBool(vars.VMBool{NID: n.XID, PID: pid}, val)
}
func (n *NodeBase) UnsetBool(pid bools.BoolVar) {
	n.XScene.VM().UnsetBool(vars.VMBool{NID: n.XID, PID: pid})
}

func (n *NodeBase) Int(pid ints.IntVar) (int, bool) {
	return n.XScene.VM().Int(vars.VMInt{NID: n.XID, PID: pid})
}
func (n *NodeBase) SetInt(pid ints.IntVar, val int) {
	n.XScene.VM().SetInt(vars.VMInt{NID: n.XID, PID: pid}, val)
}
func (n *NodeBase) UnsetInt(pid ints.IntVar) {
	n.XScene.VM().UnsetInt(vars.VMInt{NID: n.XID, PID: pid})
}

func (n *NodeBase) Float(pid floats.FloatVar) (float64, bool) {
	return n.XScene.VM().Float(vars.VMFloat{NID: n.XID, PID: pid})
}
func (n *NodeBase) SetFloat(pid floats.FloatVar, val float64) {
	n.XScene.VM().SetFloat(vars.VMFloat{NID: n.XID, PID: pid}, val)
}
func (n *NodeBase) UnsetFloat(pid floats.FloatVar) {
	n.XScene.VM().UnsetFloat(vars.VMFloat{NID: n.XID, PID: pid})
}

func (n *NodeBase) Str(pid strs.StrVar) (string, bool) {
	return n.XScene.VM().Str(vars.VMStr{NID: n.XID, PID: pid})
}
func (n *NodeBase) SetStr(pid strs.StrVar, val string) {
	n.XScene.VM().SetStr(vars.VMStr{NID: n.XID, PID: pid}, val)
}
func (n *NodeBase) UnsetStr(pid strs.StrVar) {
	n.XScene.VM().UnsetStr(vars.VMStr{NID: n.XID, PID: pid})
}

func (n *NodeBase) Rect(pid rects.RectVar) (image.Rectangle, bool) {
	return n.XScene.VM().Rect(vars.VMRect{NID: n.XID, PID: pid})
}
func (n *NodeBase) SetRect(pid rects.RectVar, val image.Rectangle) {
	n.XScene.VM().SetRect(vars.VMRect{NID: n.XID, PID: pid}, val)
}
func (n *NodeBase) UnsetRect(pid rects.RectVar) {
	n.XScene.VM().UnsetRect(vars.VMRect{NID: n.XID, PID: pid})
}

func (n *NodeBase) Img(pid imgs.ImgVar) (*ebiten.Image, bool) {
	return n.XScene.VM().Img(vars.VMImg{NID: n.XID, PID: pid})
}
func (n *NodeBase) SetImg(pid imgs.ImgVar, val *ebiten.Image) {
	n.XScene.VM().SetImg(vars.VMImg{NID: n.XID, PID: pid}, val)
}
func (n *NodeBase) UnsetImg(pid imgs.ImgVar) {
	n.XScene.VM().UnsetImg(vars.VMImg{NID: n.XID, PID: pid})
}

func MakeNodeBase(
	scene arcadegame.Scene,
) *NodeBase {
	id := uniqueid.UniqueID()
	n := &NodeBase{
		XScene: scene,
		XID:    id,
	}
	return n
}
