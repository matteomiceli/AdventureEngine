package command

import (
	"adventureengine/content"
	"adventureengine/pkg/color"
	"adventureengine/pkg/game"
	"adventureengine/pkg/help"
	"adventureengine/pkg/inventory"
	"adventureengine/state"
	"fmt"
	"strings"
)

func Input() [2]string {
	var cmd string
	var subject string
	drawPrompt()
	fmt.Scanln(&cmd, &subject)

	return [2]string{cmd, subject}
}

func CommandController(input [2]string) {
	cmd := strings.ToLower(input[0])
	subject := strings.ToLower(input[1])

	switch cmd {

	case "inventory":
		inventory.Draw()

	case "take":
		if subject == "" {
			fmt.Printf("The %s command requires a subject \n", color.PaintText(color.Yellow, "TAKE"))
			break
		}
		inventory.Add(subject)

	case "drop":
		inventory.Drop(subject)

	case "help":
		help.Lookup(subject)

	case "walk":
		game.Walk(subject)
	}

}

func drawPrompt() {
	fmt.Println()
	fmt.Printf(color.PaintText(color.Cyan, "[%s]%s%s %s "), content.Locations[state.CurrentLocationId].Display, drawHealth(), drawShields(), color.PaintText(color.Yellow, ">>"))
}

func drawHealth() string {
	healthIcon := ""
	for i := 0; i < state.Health; i++ {
		healthIcon = healthIcon + "♥"
	}
	for i := 0; i < 3-state.Health; i++ {
		healthIcon = healthIcon + "♡"
	}
	return color.PaintText(color.Red, healthIcon)
}

func drawShields() string {
	shieldIcon := ""
	for i := 0; i < state.Shield; i++ {
		shieldIcon = shieldIcon + "⛨"
	}
	return shieldIcon
}
