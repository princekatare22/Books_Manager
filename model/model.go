package model

type Book struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price   int    `json:"price"`
	Genre string `json:"genre"`
}