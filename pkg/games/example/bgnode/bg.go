package bgnode

import (
	"image"
	"image/color"
	"math/rand/v2"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/nodes"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/bools"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/imgs"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/ints"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/rects"
	"github.com/hajimehoshi/ebiten/v2"
)

var (
	ObstacleImage = (func() *ebiten.Image {
		img := ebiten.NewImage(16, 16)
		img.Fill(color.NRGBA{255, 0, 0, 255})
		return img
	})()
	BGImage = (func() *ebiten.Image {
		img := ebiten.NewImage(16, 16)
		img.Fill(color.NRGBA{0, 200, 60, 255})
		return img
	})()
	InitFunc = func(n nodes.Node) {
		n.SetBool(bools.IsActive, true)
		n.SetBool(bools.IsVisible, true)
		n.SetInt(ints.X, rand.IntN(25)*16)
		n.SetInt(ints.Y, rand.IntN(25)*16)
		n.SetImg(imgs.Primary, BGImage)

		if rand.IntN(1) == 1 {
			n.SetRect(rects.Obstacle, image.Rect(0, 0, 16, 16))
		}
	}
	DeleteFunc = func(n nodes.Node) {
		n.UnsetBool(bools.IsActive)
		n.UnsetBool(bools.IsVisible)
		n.UnsetInt(ints.X)
		n.UnsetInt(ints.Y)
		n.UnsetRect(rects.Obstacle)
		n.UnsetImg(imgs.Primary)
	}
	DrawFunc = func(n nodes.Node, sc *ebiten.Image) {
		if isActive, _ := n.Bool(bools.IsActive); !isActive {
			return
		}
		if isVisible, _ := n.Bool(bools.IsVisible); !isVisible {
			return
		}

		x, _ := n.Int(ints.X)
		y, _ := n.Int(ints.Y)

		if img, ok := n.Img(imgs.Primary); ok {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			sc.DrawImage(img, op)
		}

		if _, ok := n.Rect(rects.Obstacle); ok {
			op := &ebiten.DrawImageOptions{}
			op.GeoM.Translate(float64(x), float64(y))
			sc.DrawImage(ObstacleImage, op)
		}
	}
	UpdateFunc = func(n nodes.Node) {
		if isActive, _ := n.Bool(bools.IsActive); !isActive {
			return
		}

		dx := rand.IntN(3) - 1
		x, _ := n.Int(ints.X)
		n.SetInt(ints.X, x+dx)

		dy := rand.IntN(3) - 1
		y, _ := n.Int(ints.Y)
		n.SetInt(ints.Y, y+dy)
	}
)

func MakeBG(s arcadegame.Scene) *nodes.NodeBase {
	n := nodes.MakeNodeBase(s)
	n.XInitFunc = &InitFunc
	n.XDeleteFunc = &DeleteFunc
	n.XDrawFunc = &DrawFunc
	n.XUpdateFunc = &UpdateFunc
	return n
}
