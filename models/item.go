package models

type Edible struct {
	CanEat bool
}

type Item struct {
	Id         string
	Consumable bool
	Health     int
}
