package sheets

import (
	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/settings"
)

const (
	// Row 1
	Empty000 arcadegame.SpriteSheetFrameIndex = iota
	GrassLeft
	GrassCenter
	GrassRight
	Empty001

	// Row 2
	CloudLeft
	CloudCenter
	CloudRight
	SteelCenter
	SteelRepeat
)

func MakeJollySheet() (arcadegame.SpriteSheet, error) {
	sheet := arcadegame.SpriteSheet{
		Frames: arcadegame.AllocateFrameRects(5, 2, settings.TileSize, settings.TileSize),
	}
	img, err := arcadegame.LoadEbitenPNG("assets/tiles.png")
	if err != nil {
		return sheet, err
	}
	sheet.Init(img)
	return sheet, nil
}
