package arcadegame

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type NodeTypeToken = uint8

type NPNode struct {
	// 4 byte
	ID   UID
	Type NodeTypeToken
}

type NPPos struct {
	XPos float32
	YPos float32
}

type NPVel struct {
	XVel float32
	YVel float32
}

type NPImg struct {
	Img *ebiten.Image
}

type NPFlags struct {
	Discarded bool
	CanUpdate bool
	CanDraw   bool
}

type NPSensor struct {
	Sensor Sensor
}

type Node2D struct {
	NPNode
	NPPos
	NPImg
	NPFlags
}
