package content

import (
	"adventureengine/model"
)

var Items = map[string]*model.Item{
	"lighter": {
		Id: "lighter",
	},
	"potion": {
		Id:         "potion",
		Consumable: true,
		Health:     2,
		Display:    "Health Potion",
	},
	"key": {
		Id:      "key",
		Display: "Closet Key",
	},
}
