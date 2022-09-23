package inventory

import (
	"adventureengine/content"
	"adventureengine/helpers"
	"adventureengine/pkg/color"
	"adventureengine/pkg/game"
	"adventureengine/pkg/player"
	"adventureengine/state"
	"fmt"
	"strings"
)

func Draw() {
	for _, element := range state.Store {
		if element == "" {
			break
		}
		fmt.Println(element)
	}
}

func Add(object string) {
	if !helpers.LocationContainsItem(object, game.CurrentLocation().Items) {
		fmt.Printf("%s not found nearby. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
		return
	}
	isFull := true
	for i, element := range state.Store {
		if element == object {
			fmt.Printf("You're already carrying a %s . \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			return
		}
		if element == "" {
			state.Store[i] = object
			isFull = false
			fmt.Printf("%s added to inventory! \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			return
		}
	}
	if isFull {
		fmt.Printf("Inventory full, please %s an item before trying to pick up something else. \n", color.PaintText(color.Yellow, "DROP"))
	}
}

func Drop(object string) {
	if object == "" {
		fmt.Printf("Select an object from your inventory to %s \n", color.PaintText(color.Yellow, strings.ToUpper("DROP")))
		return
	}
	for i, element := range state.Store {
		if element == object {
			fmt.Printf("%s dropped. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			state.Store[i] = ""
			return
		}
	}
	fmt.Printf("You are not carrying a %s. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
}

func Consume(object string) {
	if object == "" {
		fmt.Printf("Select an object from your inventory to %s \n", color.PaintText(color.Yellow, strings.ToUpper("consume")))
		return
	}

	for i, element := range state.Store {
		if element == object {
			foundItem := content.Items[state.Store[i]]

			if !foundItem.Consumable {
				fmt.Printf("%s cannot be consumed. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
				return
			}
			player.IncrementHealth(foundItem.Health)
			fmt.Printf("%s consumed. %s %v health. \n", color.PaintText(color.Yellow, strings.ToUpper(object)), normalizeIntValueToHumanReadable(foundItem.Health), foundItem.Health)
			state.Store[i] = ""
			return
		}
	}

	fmt.Printf("You are not carrying a %s. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
}

func normalizeIntValueToHumanReadable(val int) string {
	if val > 0 {
		return "Gained"
	} else if val < 0 {
		return "Lost"
	}
	return ""
}
