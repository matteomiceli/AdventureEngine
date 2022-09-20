package player

import "adventureengine/state"

func IncrementHealth() {
	state.Health++
}

func DecrementHealth() {
	state.Health--
}
