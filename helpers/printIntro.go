package helpers

import (
	"adventureengine/pkg/game"
	"fmt"
)

func PrintIntro() {
	fmt.Println(game.CurrentLocation().Message)
}
