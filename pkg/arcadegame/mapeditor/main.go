package mapeditor

import (
	"errors"
	"flag"
	"log"
	"os"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/hajimehoshi/ebiten/v2"
)

/*
 * This editor is DUMB and only allows hardcoded tilesets and painting for now.
 * Only barely easier than []tiles{0,0,0,0...} allocation
 */

const (
	FlagMap          = "map"
	FlagSheet        = "sheet"
	FlagMapWidth     = "mapwidth"
	FlagMapHeight    = "mapheight"
	FlagTileWidthPx  = "tilewidth"
	FlagTileHeightPx = "tileheight"
)

var CmdError = errors.New("expected 'new' 'edit' or 'settings' subcommand")

type FlagSettings struct {
	Map          string
	Sheet        string
	MapWidth     uint
	MapHeight    uint
	TileWidthPx  uint
	TileHeightPx uint
	IsNew        bool
}

func ParseFlags() (FlagSettings, error) {
	s := FlagSettings{
		Map:          *flag.String(FlagMap, "", "path to map file"),
		Sheet:        *flag.String(FlagSheet, "", "path to sprite sheet png"),
		MapWidth:     *flag.Uint(FlagMapWidth, 0, "width of map in tile units"),
		MapHeight:    *flag.Uint(FlagMapHeight, 0, "height of map in tile units"),
		TileWidthPx:  *flag.Uint(FlagMapHeight, 0, "width of sheet tiles in pixels"),
		TileHeightPx: *flag.Uint(FlagMapHeight, 0, "height of sheet tiles in pixels"),
	}
	if s.Map == "" {
		return s, errors.New("Map Path is empty")
	}
	if s.Sheet == "" {
		return s, errors.New("Sheet Path is empty")
	}
	if s.MapWidth == 0 {
		return s, errors.New("Map Width Units is empty")
	}
	if s.MapHeight == 0 {
		return s, errors.New("Map Height Units is empty")
	}
	if s.TileWidthPx == 0 {
		return s, errors.New("Tile Width Pixels is empty")
	}
	if s.TileHeightPx == 0 {
		return s, errors.New("Tile Height Pixels is empty")
	}
	return s, nil
}

func Main() error {
	st, err := ParseFlags()
	if err != nil {
		return err
	}

	// Load Sheet
	si, err := arcadegame.LoadEbitenPNG(st.Sheet)
	if err != nil {
		return errors.New("Unable to open Sprite Sheet from path")
	}
	if si.Bounds().Dx()%int(st.TileWidthPx) != 0 {
		return errors.New("Unexpected remainder on Sheet width vs tile width pixels")
	}
	if si.Bounds().Dy()%int(st.TileHeightPx) != 0 {
		return errors.New("Unexpected remainder on Sheet width vs tile width pixels")
	}

	// Load map
	mfb, err := os.ReadFile(st.Map)
	if err != nil {
		st.IsNew = true
		mfb = make([]byte, int(st.MapWidth*st.MapHeight))
	}
	if len(mfb) != int(st.MapWidth*st.MapHeight) {
		return errors.New("Unexpected map size with given width and height")
	}

	ss := arcadegame.SpriteSheet{
		Frames: arcadegame.AllocateFrameRects(int(st.MapWidth), int(st.MapHeight), int(st.TileWidthPx), int(st.TileHeightPx)),
	}
	ss.Init(si)

	tm := arcadegame.TileMap{
		TileSizePx:  uint8(st.TileWidthPx),
		MapWidth:    uint16(st.MapWidth),
		SpriteSheet: &ss,
	}
	_ = tm
	return nil
}

func CmdEdit(mapPath string) error {

	ebiten.SetWindowSize(WindowWidth, WindowHeight)
	ebiten.SetWindowTitle(WindowTitle)
	ebiten.SetTPS(TPS)
	ebiten.SetVsyncEnabled(VSyncEnabled)

	g := &Game{}
	g.Init()
	g.Enter()

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
	return nil
}
