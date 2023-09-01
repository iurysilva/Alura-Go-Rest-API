package models

type Personality struct {
	id      int
	Name    string `json:"name"`
	History string `json:"history"`
}
