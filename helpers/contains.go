package helpers

import (
	"adventureengine/model"
	"fmt"
)

func LocationContainsItem(object string, items []model.Item) bool {
	for i, item := range items {
		fmt.Print(i, item)
		if object == item.Id {
			return true
		}
	}
	return false
}
