package models

type Book struct {
	ID    uint    `json:"id"`
	Isbn  int     `json:"isbn"`
	Title string  `json:"title"`
	Price float32 `json:"price"`
}
