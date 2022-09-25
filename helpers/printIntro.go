package helpers

import (
	"adventureengine/pkg/location"
	"fmt"
)

func PrintIntro() {
	fmt.Println(location.CurrentLocation().Message)
}
