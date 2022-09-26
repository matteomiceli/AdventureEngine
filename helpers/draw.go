package helpers

import "fmt"

func DrawItems(items []string) {
	for _, element := range items {
		if element == "" {
			break
		}
		fmt.Println(element)
	}
}
