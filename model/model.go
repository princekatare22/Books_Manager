package model

type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price   int    `json:"price"`
	Genre string `json:"genre"`
}