package editor

import (
	"fmt"
	"os"

	"github.com/BossRighteous/arcade-games/pkg/arcadegame/mapeditor"
)

func main() {
	err := mapeditor.Main()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
