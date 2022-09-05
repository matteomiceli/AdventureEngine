package command

import (
	"fmt"
	"saran/pkg/color"
	"saran/pkg/inventory"
	"strings"
)

func Input() [2]string {
	var cmd string
	var subject string
	fmt.Print(color.PaintText(color.Cyan, "[Dark Cellar] --> "))
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
		if subject != "" {
			// add to inventory
			fmt.Printf("%s has been added to your inventory \n", color.PaintText(color.Yellow, subject))
		}
	}
}
