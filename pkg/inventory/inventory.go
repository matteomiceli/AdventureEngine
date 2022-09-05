package inventory

import "fmt"

type Item struct {
	name string
}

var Store = make(map[string]Item)

func Add(item Item) {
	//
}

func Draw() {
	Store["Hack"] = Item{name: "Hack"}
	Store["Test"] = Item{name: "Test"}
	for _, element := range Store {
		fmt.Println(element.name)
	}
}
