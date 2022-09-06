package inventory

import (
	"adventureengine/pkg/color"
	"fmt"
)

type Item struct {
	Name string
}

var Store = map[string]Item{
	"hack": {Name: "Hack"},
	"name": {Name: "Test"},
	"":     {Name: ""},
}

func Draw() {
	for _, element := range Store {
		fmt.Println(element.Name)
	}
}

func Add(object string) {
	if len(Store) >= 3 {
		fmt.Printf("Inventory full, please %s an item before trying to pick up something else. \n", color.PaintText(color.Yellow, "DROP"))
	}
	// if object does not exist in list of items in room, return

	for _, element := range Store {
		if element.Name != "" {

		}

		fmt.Println(element)
	}
}
