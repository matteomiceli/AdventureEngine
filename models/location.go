package models

type Location struct {
	Id      string
	Message string
	Events  []Event
	GoTo    GoTo
	Items   []Item
}

type GoTo map[string]string
