package utils

import (
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/settings"
	"github.com/hajimehoshi/ebiten/v2"
)

func IsAnyKeyPressed() bool {
	keys := []ebiten.Key{
		settings.P1Coin,
		settings.P1Start,
		settings.P1Right,
		settings.P1Left,
		settings.P1Up,
		settings.P1Down,
		settings.P1B1,
		settings.P1B2,
		settings.P1B3,
		settings.P1B4,
		settings.P1B5,
		settings.P1B6,
		settings.P1B7,
		settings.P1B8,
		settings.P2Coin,
		settings.P2Start,
		settings.P2Right,
		settings.P2Left,
		settings.P2Up,
		settings.P2Down,
		settings.P2B1,
		settings.P2B2,
		settings.P2B3,
		settings.P2B4,
		settings.P2B5,
		settings.P2B6,
		settings.P2B7,
		settings.P2B8,
	}
	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			return true
		}
	}
	return false
}
