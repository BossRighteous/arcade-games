package arcadegame

import (
	"image"

	"github.com/hajimehoshi/ebiten/v2"
)

type AnimationFrame struct {
	SpriteSheetRect image.Rectangle
	FrameHold       int
}

type Animation struct {
	ID     int
	Frames []AnimationFrame
}

type Sprite struct {
	Sheet      *ebiten.Image
	Animations []Animation
}

type SpriteInstance struct {
	sp        *Sprite
	currAnim  int
	currFrame int
	isPlaying bool
}

type SpriteSheetToken = uint8
type SpriteSheetFrameIndex = uint16 // Index of frame for naming
type SpriteSheet struct {
	Img       *ebiten.Image
	Frames    []image.Rectangle
	FrameImgs []*ebiten.Image
}

func (ss *SpriteSheet) Init(img *ebiten.Image) {
	ss.Img = img
	ss.FrameImgs = make([]*ebiten.Image, len(ss.Frames))
	for fi, fr := range ss.Frames {
		ss.FrameImgs[fi] = ss.Img.SubImage(fr).(*ebiten.Image)
	}
}

func AllocateFrameRects(tw int, th int, rectW int, rectH int) []image.Rectangle {
	rs := make([]image.Rectangle, (tw * th))
	i := 0
	for ty := range th {
		for tx := range tw {
			rs[i] = image.Rectangle{
				Min: image.Point{X: tx * rectW, Y: ty * rectH},
				Max: image.Point{X: (tx * rectW) + rectW, Y: (ty * rectH) + rectH},
			}
			i++
		}
	}
	return rs
}
