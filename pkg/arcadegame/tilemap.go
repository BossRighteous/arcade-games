package arcadegame

import "errors"

// This type may be marshalable and should remain simple

// Used to indicate runtime factories for on-screen content
type TileToken int

// Used to indicate scene factory for Sheet Image load
type TileSheetToken int

type TileMap struct {
	TileSizePx int
	MapWidth   int //number of tiles per 'row'
	SheetToken TileSheetToken
	Tiles      []TileToken
}

var TileOutOfRangeError = errors.New("Tile index out of range")

func (m *TileMap) GetTileAt(tx int, ty int) (TileToken, error) {
	ti := (ty * m.MapWidth) + tx
	if ti >= len(m.Tiles) || ti < 0 {
		return 0, TileOutOfRangeError
	}
	return m.Tiles[ti], nil
}

func (m *TileMap) GetRowFrom(tx int, ty int, w int) ([]TileToken, error) {
	return m.GetTileRect(tx, ty, w, 1)
}

func (m *TileMap) GetColumnFrom(tx int, ty int, h int) ([]TileToken, error) {
	return m.GetTileRect(tx, ty, 1, h)
}

func (m *TileMap) GetTileRect(tx int, ty int, w int, h int) ([]TileToken, error) {
	var tiles []TileToken
	if w <= 0 || tx+w > m.MapWidth || h <= 0 || ty+h > int(len(m.Tiles)/m.MapWidth) {
		return tiles, TileOutOfRangeError
	}
	tiles = make([]TileToken, w*h)
	ti := (ty * m.MapWidth) + tx
	i := 0
	for hi := range h {
		ci := ti + (hi * m.MapWidth)
		if ci+w > len(m.Tiles) {
			return []TileToken{}, TileOutOfRangeError
		}
		rt := m.Tiles[ci : ci+w]
		for _, t := range rt {
			tiles[i] = t
			i++
		}
	}
	return tiles, nil
}
