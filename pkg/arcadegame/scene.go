package arcadegame

import (
	"github.com/BossRighteous/arcade-games/pkg/arcadegame/vars/vm"
	"github.com/hajimehoshi/ebiten/v2"
)

type Scene interface {
	Init(GameRoot)
	Enter()
	Update()
	Draw(screen *ebiten.Image)
	Exit()
	Root() GameRoot
	VM() vm.VM
}
