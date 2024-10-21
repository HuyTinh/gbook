package response

import "GBook_be/internal/models"

type BookResponse struct {
	ID     uint          `json:"id"`
	Title  string        `json:"title"`
	Genre  string        `json:"genre"` // Foreign key to Author
	Author models.Author `json:"author"`
}
