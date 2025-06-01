package arcadegame

import (
	"encoding/binary"
	"errors"
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

// This type may be marshalable and should remain simple

var P0 = image.Point{X: 0, Y: 0}

type TileMap struct {
	TileSizePx  uint8
	MapWidth    uint16 //number of tiles per 'row'
	SpriteSheet *SpriteSheet
	Tiles       []SpriteSheetFrameIndex
}

type Tile struct {
	SheetIndex SpriteSheetFrameIndex
	TX         uint16
	TY         uint16
}

var TileOutOfRangeError = errors.New("Tile index out of range")

func (m *TileMap) GetTileAt(tx uint16, ty uint16) (Tile, error) {
	ti := (ty * m.MapWidth) + tx
	if int(ti) >= len(m.Tiles) || ti < 0 {
		return Tile{}, TileOutOfRangeError
	}
	return Tile{
		SheetIndex: m.Tiles[ti],
		TX:         tx,
		TY:         ty,
	}, nil
}

func (m *TileMap) GetRowFrom(tx uint16, ty uint16, w uint16) ([]Tile, error) {
	return m.GetTileRect(tx, ty, w, 1)
}

func (m *TileMap) GetColumnFrom(tx uint16, ty uint16, h uint16) ([]Tile, error) {
	return m.GetTileRect(tx, ty, 1, h)
}

func (m *TileMap) GetTileRect(tx uint16, ty uint16, w uint16, h uint16) ([]Tile, error) {
	var tiles []Tile
	if w <= 0 || tx+w > m.MapWidth || h <= 0 || int(ty+h) > len(m.Tiles)/int(m.MapWidth) {
		return tiles, TileOutOfRangeError
	}
	tiles = make([]Tile, w*h)
	ti := (ty * m.MapWidth) + tx
	i := 0
	for hi := range h {
		ci := ti + (hi * m.MapWidth)
		if int(ci+w) > len(m.Tiles) {
			return []Tile{}, TileOutOfRangeError
		}
		rt := m.Tiles[ci : ci+w]
		for rtx, t := range rt {
			tiles[i] = Tile{
				SheetIndex: t,
				TX:         uint16(rtx),
				TY:         uint16(hi + ty),
			}
			i++
		}
	}
	return tiles, nil
}

func (m *TileMap) DrawTileRect(img *ebiten.Image, o image.Point, tx uint16, ty uint16, w uint16, h uint16) error {
	ts, err := m.GetTileRect(tx, ty, w, h)
	if err != nil {
		return err
	}
	fi := m.SpriteSheet.FrameImgs
	sz := float64(m.TileSizePx)
	ox := float64(o.X)
	oy := float64(o.Y)
	for _, t := range ts {
		if t.SheetIndex == 0 {
			continue
		}
		ob := ebiten.DrawImageOptions{}
		ob.GeoM.Translate((float64(t.TX)*sz)+ox, (float64(t.TY)*sz)+oy)
		img.DrawImage(fi[t.SheetIndex], &ob)
	}
	return nil
}

func (m *TileMap) DrawTiles(img *ebiten.Image) error {
	return m.DrawTileRect(img, P0, 0, 0, m.MapWidth, uint16(len(m.Tiles)/int(m.MapWidth)))
}

func MarshalTiles(ts []SpriteSheetFrameIndex) []byte {
	// uint16 is 2 bytes
	bl := 2
	buf := make([]byte, len(ts)*bl)
	for i, t := range ts {
		binary.LittleEndian.PutUint16(buf[(i*bl):(i*bl)+bl], t)
	}
	return buf
}

func UnmarshalTiles(buf []byte) []SpriteSheetFrameIndex {
	// uint16 is 2 bytes
	bl := 2
	l := len(buf) / bl
	ts := make([]SpriteSheetFrameIndex, l)
	for i := range ts {
		ts[i] = binary.LittleEndian.Uint16(buf[i*bl : (i*bl)+bl])
	}
	return ts
}
