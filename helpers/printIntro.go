package helpers

import (
	"adventureengine/state"
	"fmt"
)

func PrintIntro() {
	fmt.Println(state.CurrentLocation().Message)
}
