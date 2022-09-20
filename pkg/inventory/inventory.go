package inventory

import (
	"adventureengine/content"
	"adventureengine/helpers"
	"adventureengine/pkg/color"
	"adventureengine/pkg/game"
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
		fmt.Printf("%s not found in room. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
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
	found := false
	for i, element := range state.Store {
		if element == object {
			found = true
			fmt.Printf("%s dropped. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
			state.Store[i] = ""
			return
		}
	}
	if !found {
		fmt.Printf("You are not carrying a %s. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
	}
}

func Eat(object string) {
	if object == "" {
		fmt.Printf("Select an object from your inventory to %s \n", color.PaintText(color.Yellow, strings.ToUpper("EAT")))
		return
	}

	found := false
	for i, element := range state.Store {
		if element == object {
			found = true
			foundItem := content.Items[state.Store[i]]
			if foundItem.Eat != nil {

				fmt.Printf("%s consumed. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
				state.Store[i] = ""
				return
			}
		}
	}
	if !found {
		fmt.Printf("You are not carrying a %s. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
	}

}
