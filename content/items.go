package content

import (
	"adventureengine/models"
)

var Items = map[string]models.Item{
	"lighter": {
		Id: "lighter",
	},
	"potion": {
		Id:         "potion",
		Consumable: true,
		Health:     2,
	},
}
