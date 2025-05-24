package node2d

import (
	"fmt"
	"image"
	"image/color"
	"math"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/hajimehoshi/ebiten/v2"
)

type Node2DProps struct {
	IsVisible     bool
	IsActive      bool
	XPos          float64
	YPos          float64
	XVel          float64
	YVel          float64
	CollisionMask byte
	BoundingRect  image.Rectangle
	Sensors       []image.Rectangle
	Sprite        arcadegame.SpriteInstance
	SpriteOffsetX int
	SpriteOffsetY int
}

var Node2DDrawFunc = func(n arcadegame.Node, s *ebiten.Image) {
	p, ok := n.Props().(*Node2DProps)
	if !ok || !p.IsActive || !p.IsVisible {
		return
	}
	img, err := p.Sprite.DrawFrame()
	if err != nil {
		fmt.Println(err)
		return
	}

	tx := math.Floor(p.XPos + float64(p.SpriteOffsetX))
	ty := math.Floor(p.YPos + float64(p.SpriteOffsetY))

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(tx, ty)
	s.DrawImage(img, op)
}

var TmpBGSprite = (func() arcadegame.Sprite {
	img := ebiten.NewImage(16, 16)
	img.Fill(color.NRGBA{0, 200, 60, 255})
	sp := arcadegame.Sprite{
		Sheet: img,
	}
	return sp
})()
