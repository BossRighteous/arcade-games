package vars

import (
	"image"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame/uniqueid"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/bools"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/floats"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/imgs"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/ints"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/rects"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/strs"
	"github.com/hajimehoshi/ebiten/v2"
)

// Maybe I should have done a ECS? Oh well

type VMType int

const (
	TypeBool VMType = iota
	TypeInt
	TypeFloat
	TypeStr
	TypeRect
	TypeImg
)

type VMI struct {
	NID uniqueid.UID // Node UniqueID
	PID int          // iota of property
	TID VMType
}

type VMBool struct {
	NID uniqueid.UID
	PID bools.BoolVar
}

type VMInt struct {
	NID uniqueid.UID
	PID ints.IntVar
}

type VMFloat struct {
	NID uniqueid.UID
	PID floats.FloatVar
}

type VMStr struct {
	NID uniqueid.UID
	PID strs.StrVar
}

type VMRect struct {
	NID uniqueid.UID
	PID rects.RectVar
}

type VMImg struct {
	NID uniqueid.UID
	PID imgs.ImgVar
}

// Implements VM
type VarMap struct {
	BoolMap  map[VMBool]bool
	IntMap   map[VMInt]int
	FloatMap map[VMFloat]float64
	StrMap   map[VMStr]string
	RectMap  map[VMRect]image.Rectangle
	ImgMap   map[VMImg]*ebiten.Image
}

func NewVarMap() *VarMap {
	return &VarMap{
		BoolMap:  make(map[VMBool]bool),
		IntMap:   make(map[VMInt]int),
		FloatMap: make(map[VMFloat]float64),
		StrMap:   make(map[VMStr]string),
		RectMap:  make(map[VMRect]image.Rectangle),
		ImgMap:   make(map[VMImg]*ebiten.Image),
	}
}

func (vm *VarMap) ClearAll() {
	for key := range vm.BoolMap {
		delete(vm.BoolMap, key)
	}
	vm.BoolMap = make(map[VMBool]bool)

	for key := range vm.IntMap {
		delete(vm.IntMap, key)
	}
	vm.IntMap = make(map[VMInt]int)

	for key := range vm.FloatMap {
		delete(vm.FloatMap, key)
	}
	vm.FloatMap = make(map[VMFloat]float64)

	for key := range vm.StrMap {
		delete(vm.StrMap, key)
	}
	vm.StrMap = make(map[VMStr]string)

	for key := range vm.RectMap {
		delete(vm.RectMap, key)
	}
	vm.RectMap = make(map[VMRect]image.Rectangle)

	for key := range vm.ImgMap {
		delete(vm.ImgMap, key)
	}
	vm.ImgMap = make(map[VMImg]*ebiten.Image)
}

func (vm *VarMap) Bool(vmi VMBool) (bool, bool) {
	val, ok := vm.BoolMap[vmi]
	return val, ok
}
func (vm *VarMap) SetBool(vmi VMBool, val bool) {
	vm.BoolMap[vmi] = val
}
func (vm *VarMap) UnsetBool(vmi VMBool) {
	delete(vm.BoolMap, vmi)
}

func (vm *VarMap) Int(vmi VMInt) (int, bool) {
	val, ok := vm.IntMap[vmi]
	return val, ok
}
func (vm *VarMap) SetInt(vmi VMInt, val int) {
	vm.IntMap[vmi] = val
}
func (vm *VarMap) UnsetInt(vmi VMInt) {
	delete(vm.IntMap, vmi)
}

func (vm *VarMap) Float(vmi VMFloat) (float64, bool) {
	val, ok := vm.FloatMap[vmi]
	return val, ok
}
func (vm *VarMap) SetFloat(vmi VMFloat, val float64) {
	vm.FloatMap[vmi] = val
}
func (vm *VarMap) UnsetFloat(vmi VMFloat) {
	delete(vm.FloatMap, vmi)
}

func (vm *VarMap) Str(vmi VMStr) (string, bool) {
	val, ok := vm.StrMap[vmi]
	return val, ok
}
func (vm *VarMap) SetStr(vmi VMStr, val string) {
	vm.StrMap[vmi] = val
}
func (vm *VarMap) UnsetStr(vmi VMStr) {
	delete(vm.StrMap, vmi)
}

func (vm *VarMap) Rect(vmi VMRect) (image.Rectangle, bool) {
	val, ok := vm.RectMap[vmi]
	return val, ok
}
func (vm *VarMap) SetRect(vmi VMRect, val image.Rectangle) {
	vm.RectMap[vmi] = val
}
func (vm *VarMap) UnsetRect(vmi VMRect) {
	delete(vm.RectMap, vmi)
}

func (vm *VarMap) Img(vmi VMImg) (*ebiten.Image, bool) {
	val, ok := vm.ImgMap[vmi]
	return val, ok
}
func (vm *VarMap) SetImg(vmi VMImg, val *ebiten.Image) {
	vm.ImgMap[vmi] = val
}
func (vm *VarMap) UnsetImg(vmi VMImg) {
	delete(vm.ImgMap, vmi)
}

type VarMappable interface {
	Bool(bools.BoolVar) (bool, bool)
	SetBool(bools.BoolVar, bool)
	UnsetBool(bools.BoolVar)

	Int(ints.IntVar) (int, bool)
	SetInt(ints.IntVar, int)
	UnsetInt(ints.IntVar)

	Float(floats.FloatVar) (float64, bool)
	SetFloat(floats.FloatVar, float64)
	UnsetFloat(floats.FloatVar)

	Str(strs.StrVar) (string, bool)
	SetStr(strs.StrVar, string)
	UnsetStr(strs.StrVar)

	Rect(rects.RectVar) (image.Rectangle, bool)
	SetRect(rects.RectVar, image.Rectangle)
	UnsetRect(rects.RectVar)

	Img(imgs.ImgVar) (*ebiten.Image, bool)
	SetImg(imgs.ImgVar, *ebiten.Image)
	UnsetImg(imgs.ImgVar)
}
