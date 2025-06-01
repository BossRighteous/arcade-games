package nodes

import "github.com/BossRighteous/arcade-games/pkg/arcadegame"

type Player struct {
	arcadegame.Node2D
	arcadegame.NPSensor
	IsGrounded    bool
	IsClippingOff bool
}
