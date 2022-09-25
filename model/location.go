package model

type Location struct {
	Id      string
	Display string
	Message string
	Events  []Event
	GoTo    GoTo
	Items   []Item
}

type GoTo map[string]string
