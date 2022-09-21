package player

import "adventureengine/state"

func IncrementHealth(val int) {
	state.Health = state.Health + val
}

func DecrementHealth(val int) {
	state.Health = state.Health - val
}

func IncrementShields(val int) {
	state.Shield = state.Shield + val
}

func DecrementShields(val int) {
	state.Shield = state.Shield - val
}
