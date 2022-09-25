package model

type Event struct {
	cmd    string
	effect func(subject string)
}
