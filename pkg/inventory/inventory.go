package inventory

import (
	"adventureengine/content"
	"adventureengine/helpers"
	"adventureengine/pkg/color"
	"adventureengine/pkg/player"
	"adventureengine/state"
	"fmt"
	"strings"
)

func Add(object string) {
	if object == "" {
		fmt.Printf("The %s command requires a subject \n", color.PaintText(color.Yellow, "TAKE"))
		return
	}
	if !helpers.LocationContainsItem(object, state.CurrentLocation().Items) {
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
			RemoveItemFromLocation(state.CurrentLocation().Id, object)
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
			AddItemToLocation(state.CurrentLocation().Id, object)
			return
		}
	}
	fmt.Printf("You are not carrying a %s. \n", color.PaintText(color.Yellow, strings.ToUpper(object)))
}

func AddItemToLocation(locationID string, object string) {
	items := content.Locations[locationID].Items
	content.Locations[locationID].Items = append(items, object)
}

func RemoveItemFromLocation(locationID string, object string) {
	itemsInLocation := content.Locations[locationID].Items
	var itemIndex int
	for i, item := range itemsInLocation {
		if item == object {
			itemIndex = i
		}
	}

	content.Locations[locationID].Items = append(itemsInLocation[:itemIndex], itemsInLocation[itemIndex+1:]...)
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
