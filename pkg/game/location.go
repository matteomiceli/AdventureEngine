package game

type Location struct {
	name    string
	message string
	events  []Event
}

var CurrentLocation Location = Location{
	name:    "start",
	message: "You arrive at the start of your journey!",
	events: []Event{
		{cmd: "walk", to: "sanctum"},
	},
}
