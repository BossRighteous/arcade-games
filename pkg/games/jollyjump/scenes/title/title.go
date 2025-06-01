package title

import (
	"image"
	"image/color"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/settings"
	"github.com/BossRighteous/arcade-games/pkg/games/jollyjump/sheets"
	"github.com/hajimehoshi/ebiten/v2"
)

type SceneTitle struct {
	r       arcadegame.GameRoot
	bgColor color.Color
	idleF   int
	bgImg   *ebiten.Image
	m       arcadegame.TileMap
	sens    []arcadegame.Sensor
}

func (s *SceneTitle) Init(r arcadegame.GameRoot) {
	s.r = r
	s.bgColor = color.RGBA{97, 162, 255, 255}
	s.idleF = 0

	sheet, _ := sheets.MakeJollySheet()
	s.m = makeTileMap(&sheet)

	s.sens = makeSensors()

	s.bgImg = ebiten.NewImageWithOptions(
		image.Rect(0, 0, settings.ScreenWidth, settings.ScreenHeight),
		&ebiten.NewImageOptions{
			Unmanaged: true,
		},
	)
	s.bgImg.Fill(s.bgColor)
	err := s.m.DrawTiles(s.bgImg)
	if err != nil {
		panic(err)
	}

	if settings.Debug == true {
		sImg := ebiten.NewImageWithOptions(
			image.Rect(0, 0, 1024, 1024),
			&ebiten.NewImageOptions{
				Unmanaged: true,
			},
		)
		sImg.Fill(color.RGBA{200, 40, 40, 128})
		for _, sen := range s.sens {
			ob := ebiten.DrawImageOptions{}
			ob.GeoM.Translate(float64(sen.XMin), float64(sen.YMin))
			s.bgImg.DrawImage(
				sImg.SubImage(image.Rect(0, 0, int(sen.XMax-sen.XMin), int(sen.YMax-sen.YMin))).(*ebiten.Image),
				&ob,
			)
		}
	}
}

func (s *SceneTitle) Enter() {

}

func (s *SceneTitle) Update() error {
	/*
		s.idleF++
		if s.idleF > 1200 {
			s.Exit()
			s.Root().SetScene(scenes.SceneAttractToken)
		}

		if inpututil.IsKeyJustPressed(settings.Esc) {
			//s.Root().SetScene()
			s.Exit()
			return ebiten.Termination
		}*/
	return nil
}

func (s *SceneTitle) Draw(sc *ebiten.Image) {
	sc.DrawImage(s.bgImg, nil)
}

func (s *SceneTitle) Exit() {

}

func (s *SceneTitle) Root() arcadegame.GameRoot {
	return s.r
}

func makeSensors() []arcadegame.Sensor {
	return []arcadegame.Sensor{
		{XMin: 64, XMax: settings.ScreenWidth - 64, YMin: 67, YMax: settings.ScreenHeight - 64},
	}
}

func makeTileMap(s *arcadegame.SpriteSheet) arcadegame.TileMap {
	tiles := make([]arcadegame.SpriteSheetFrameIndex, 30*16)
	for i := range len(tiles) {
		tiles[i] = sheets.GrassCenter
	}

	return arcadegame.TileMap{
		TileSizePx:  settings.TileSize,
		MapWidth:    30,
		SpriteSheet: s,
		Tiles:       tiles,
	}
}
