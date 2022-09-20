package player

import "adventureengine/state"

func IncrementHealth() {
	state.Health++
}

func DecrementHealth() {
	state.Health--
}

func IncrementShields() {
	state.Shield++
}

func DecrementShields() {
	state.Shield--
}
