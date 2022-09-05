package game

type Event struct {
	cmd    string
	effect func(subject string)
	to     string
}
