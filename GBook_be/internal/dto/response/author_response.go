package response

import (
	"GBook_be/internal/models"
	"time"
)

type AuthorResponse struct {
	ID          uint          `json:"id"`
	Name        string        `json:"name"`
	Biography   string        `json:"biography"`
	DateOfBirth time.Time `gorm:"type:date"`     `json:"date_of_birth"`
	Nationality string        `json:"nationality"`
	Books       []models.Book `json:"books"`
}
