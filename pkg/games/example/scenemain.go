package example

import (
	"fmt"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/nodes"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/uniqueid"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars"
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/vm"
	"github.com/BossRighteous/arcade-games/pkg/games/example/bgnode"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type SceneMain struct {
	r    arcadegame.GameRoot
	vm   *vars.VarMap
	nMap map[uniqueid.UID]nodes.Node
	f    uint64
}

func (s *SceneMain) Init(r arcadegame.GameRoot) {
	s.r = r
	s.vm = vars.NewVarMap()
	s.nMap = make(map[uniqueid.UID]nodes.Node, 100)
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

func (s *SceneMain) Update() {
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
}

func (s *SceneMain) Draw(sc *ebiten.Image) {
	for _, n := range s.nMap {
		n.Draw(sc)
	}
	ebitenutil.DebugPrint(sc, fmt.Sprintf("Nodes: %v", len(s.nMap)))
}

func (s *SceneMain) Exit() {
	for _, n := range s.nMap {
		n.Exit()
	}
	s.vm.ClearAll()
}

func (s *SceneMain) Root() arcadegame.GameRoot {
	return s.r
}

func (s *SceneMain) VM() vm.VM {
	return s.vm
}
