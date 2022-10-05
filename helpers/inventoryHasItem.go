package helpers

import "adventureengine/state"

func InventoryHasItem(itemId string) bool {
	for _, item := range state.Store {
		if item == itemId {
			return true
		}
	}
	return false
}
