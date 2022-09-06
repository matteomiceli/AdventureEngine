package helper

import (
	"adventureengine/pkg/inventory"
	"fmt"
)

func LocationContainsItem(object string, items []inventory.Item) bool {
	for i, item := range items {
		fmt.Print(i, item)
		if object == item.Name {
			return true
		}
	}
	return false
}
