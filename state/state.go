package state

import "adventureengine/model"

var CurrentLocationId string = "grandEntrance"

var Store = [3]string{
	"potion",
	"",
	"",
}

var Player model.Player = model.Player{
	Health: 3,
	Shield: 1,
}
