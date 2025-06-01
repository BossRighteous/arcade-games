package arcadegame

import (
	"image"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadEbitenPNG(p string) (*ebiten.Image, error) {
	img, err := LoadPNGFromPath(p)
	if err != nil {
		return nil, err
	}
	return ImageToEbiten(img), nil
}

func LoadPNGFromPath(p string) (image.Image, error) {
	reader, err := os.Open(p)
	if err != nil {
		return nil, err
	}
	return png.Decode(reader)
}

func ImageToEbiten(img image.Image) *ebiten.Image {
	return ebiten.NewImageFromImage(img)
}
