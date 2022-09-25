package helpers

import "fmt"

func LocationContainsItem(object string, items []string) bool {
	for _, item := range items {
		if object == item {
			fmt.Print("exists")
			return true
		}
	}
	return false
}
