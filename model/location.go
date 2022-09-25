package model

type Location struct {
	Id      string
	Display string
	Message string
	Events  []Event
	GoTo    GoTo
	Items   []string
}

type GoTo map[string]string
