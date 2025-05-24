package bgnode

import (
	"math/rand/v2"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/nodes"
	"github.com/BossRighteous/arcade-games/pkg/games/example/node2d"
)

var (
	InitFunc = func(n nodes.Node) {
		p, ok := n.Props().(*node2d.Node2DProps)
		if !ok || !p.IsActive {
			return
		}

		p.XPos = float64(rand.IntN(25) * 16)
		p.YPos = float64(rand.IntN(25) * 16)
		p.XVel = (rand.Float64() * 3) - 1.5
		p.YVel = (rand.Float64() * 3) - 1.5
	}
	UpdateFunc = func(n nodes.Node) {
		p, ok := n.Props().(*node2d.Node2DProps)
		if !ok || !p.IsActive {
			return
		}
		p.XPos = p.XPos + p.XVel
		p.YPos = p.YPos + p.YVel

		wx := float64(320)
		wy := float64(240)

		if p.XPos > wx {
			p.XPos = p.XPos - wx
		} else if p.XPos < 0 {
			p.XPos = wx + p.XPos
		}

		if p.YPos > wy {
			p.YPos = p.YPos - wy
		} else if p.YPos < 0 {
			p.YPos = wy + p.YPos
		}
	}
)

func MakeBG(s arcadegame.Scene) *nodes.NodeBase {
	n := nodes.MakeNodeBase(
		s, nodes.LifecycleHooks{
			InitFunc:   &InitFunc,
			DrawFunc:   &node2d.Node2DDrawFunc,
			UpdateFunc: &UpdateFunc,
		},
		// Pass as ref!
		&node2d.Node2DProps{
			IsActive:  true,
			IsVisible: true,
			Sprite:    arcadegame.MakeSpriteInstance(&node2d.TmpBGSprite),
		},
	)
	return n
}
