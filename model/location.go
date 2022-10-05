package model

type Location struct {
	Id      string
	Display string
	Message string
	Events  []Event
	GoTo    GoTo
	Items   []string
	Door    Door
}

type GoTo map[string]string

type Door struct {
	Locked bool
	Key    string
}
