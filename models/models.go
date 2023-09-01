package models

type Personality struct {
	Id      int
	Name    string `json:"name"`
	History string `json:"history"`
}
