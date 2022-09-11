package helpers

import (
	"adventureengine/models"
	"fmt"
)

func LocationContainsItem(object string, items []models.Item) bool {
	for i, item := range items {
		fmt.Print(i, item)
		if object == item.Id {
			return true
		}
	}
	return false
}
