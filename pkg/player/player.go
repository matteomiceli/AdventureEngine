package player

import "adventureengine/state"

func IncrementHealth(val int) {
	state.Player.Health = state.Player.Health + val
}

func DecrementHealth(val int) {
	state.Player.Health = state.Player.Health - val
}

func IncrementShields(val int) {
	state.Player.Shield = state.Player.Shield + val
}

func DecrementShields(val int) {
	state.Player.Shield = state.Player.Shield - val
}
