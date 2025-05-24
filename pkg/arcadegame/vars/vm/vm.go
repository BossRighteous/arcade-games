package vm

import (
	"image"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars"
	"github.com/hajimehoshi/ebiten/v2"
)

// TODO: move to better pkg space

type VM interface {
	ClearAll()

	Bool(vars.VMBool) (bool, bool)
	SetBool(vars.VMBool, bool)
	UnsetBool(vars.VMBool)

	Int(vars.VMInt) (int, bool)
	SetInt(vars.VMInt, int)
	UnsetInt(vars.VMInt)

	Float(vars.VMFloat) (float64, bool)
	SetFloat(vars.VMFloat, float64)
	UnsetFloat(vars.VMFloat)

	Str(vars.VMStr) (string, bool)
	SetStr(vars.VMStr, string)
	UnsetStr(vars.VMStr)

	Rect(vars.VMRect) (image.Rectangle, bool)
	SetRect(vars.VMRect, image.Rectangle)
	UnsetRect(vars.VMRect)

	Img(vars.VMImg) (*ebiten.Image, bool)
	SetImg(vars.VMImg, *ebiten.Image)
	UnsetImg(vars.VMImg)
}
