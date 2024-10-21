package models

type Book struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Title    string `json:"title"`
	Genre    string `json:"genre"`
	AuthorID uint   `json:"author_id"` // Foreign key to Author
	Author   Author `json:"author"`    // Eager loading of Author
}
