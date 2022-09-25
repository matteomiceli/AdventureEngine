package location

import (
	"adventureengine/content"
	"adventureengine/model"
	"adventureengine/state"
	"fmt"
)

func CurrentLocation() *model.Location {
	return content.Locations[state.CurrentLocationId]
}

func Walk(direction string) {
	if !isValidDirection(direction) {
		fmt.Println("Enter a valid direction.")
		return
	}

	current := CurrentLocation()

	if _, ok := current.GoTo[direction]; !ok {
		fmt.Println("There doesn't appear to be anything here.")
		return
	}
	state.CurrentLocationId = current.GoTo[direction]
	fmt.Println(CurrentLocation().Message)
}

func isValidDirection(direction string) bool {
	if direction == "left" || direction == "right" || direction == "back" || direction == "forward" {
		return true
	}
	return false
}

func DropItem(locationID string, object string) {
	items := content.Locations[locationID].Items
	content.Locations[locationID].Items = append(items, object)
}

func PickUpItem(locationID string, object string) {
	itemsInLocation := content.Locations[locationID].Items
	var itemIndex int
	for i, item := range itemsInLocation {
		if item == object {
			itemIndex = i
		}
	}

	content.Locations[locationID].Items = append(itemsInLocation[:itemIndex], itemsInLocation[itemIndex+1:]...)
}
