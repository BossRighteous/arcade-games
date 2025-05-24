package arcadegame

import (
	"errors"
	"fmt"
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

func MakeSpriteInstance(sp *Sprite) SpriteInstance {
	return SpriteInstance{sp: sp}
}

func (spi *SpriteInstance) DrawFrame() (*ebiten.Image, error) {
	if spi.sp == nil || spi.sp.Sheet == nil {
		fmt.Println("Sprite error")
		return nil, errors.New("Sprite Draw Error")
	}
	return spi.sp.Sheet, nil
	/*
		anim := spi.currAnim
		frame := spi.currFrame
		if anim < 0 || anim >= len(spi.sp.Animations) || frame < 0 || frame >= len(spi.sp.Animations[anim].Frames) {
			return nil, errors.New("error DrawFrame out of bounds")
		}
		rect := spi.sp.Animations[anim].Frames[frame].SpriteSheetRect
		return spi.sp.Sheet.SubImage(rect).(*ebiten.Image), nil
	*/
	// TODO: subimage
}

func (sp *SpriteInstance) Play() {
	sp.isPlaying = true
}

func (sp *SpriteInstance) Stop() {
	sp.isPlaying = false
}

func (sp *SpriteInstance) IsPlaying() bool {
	return sp.isPlaying
}

func (spi *SpriteInstance) SetCurrentAnimationFrame(anim int, frame int, play bool) error {
	if anim < 0 || anim >= len(spi.sp.Animations) || frame < 0 || frame >= len(spi.sp.Animations[anim].Frames) {
		anim = 0
		frame = 0
		play = false
		return errors.New("error SetCurrentAnimationFrame out of bounds")
	}
	spi.currAnim = anim
	spi.currFrame = frame
	spi.isPlaying = play
	return nil
}

func (sp *SpriteInstance) GetCurrentAnimationFrame() (int, int, bool) {
	return sp.currAnim, sp.currFrame, sp.isPlaying
}
