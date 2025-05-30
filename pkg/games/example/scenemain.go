package example

import (
	"fmt"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/games/example/bgnode"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneMain struct {
	r    arcadegame.GameRoot
	nMap map[arcadegame.UID]arcadegame.Node
	f    uint64
}

func (s *SceneMain) Init(r arcadegame.GameRoot) {
	s.r = r
	s.nMap = make(map[arcadegame.UID]arcadegame.Node, 100)
	s.f = 0
}

func (s *SceneMain) Enter() {
	cnt := 100

	for range cnt {
		n := bgnode.MakeBG(s)
		n.Init()
		n.Enter()
		s.nMap[n.ID()] = n
	}
}

func (s *SceneMain) Update() error {
	for _, n := range s.nMap {
		n.Update()
	}
	s.f++
	if s.f%4 == 0 {
		n := bgnode.MakeBG(s)
		n.Init()
		n.Enter()
		s.nMap[n.ID()] = n
	}
	return nil
}

func (s *SceneMain) Draw(sc *ebiten.Image) {
	for _, n := range s.nMap {
		n.Draw(sc)
	}

	ebitenutil.DebugPrint(sc, fmt.Sprintf("Nodes: %v, TPS: %.2f, FPS: %.2f", len(s.nMap), ebiten.ActualTPS(), ebiten.ActualFPS()))
}

func (s *SceneMain) Exit() {
	for _, n := range s.nMap {
		n.Exit()
	}
}

func (s *SceneMain) Root() arcadegame.GameRoot {
	return s.r
}
