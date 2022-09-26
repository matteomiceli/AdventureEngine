package state

import (
	"adventureengine/content"
	"adventureengine/model"
)

var CurrentLocationId string = "grandEntrance"

func CurrentLocation() *model.Location {
	return content.Locations[CurrentLocationId]
}

var Store = [3]string{
	"potion",
	"",
	"",
}

var Player model.Player = model.Player{
	Health: 3,
	Shield: 1,
}
