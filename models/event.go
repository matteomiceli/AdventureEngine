package models

type Event struct {
	cmd    string
	effect func(subject string)
}
