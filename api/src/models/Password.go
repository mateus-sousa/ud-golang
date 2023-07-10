package models

type Password struct {
	Current string `json:"current"`
	New     string `json:"new"`
}
