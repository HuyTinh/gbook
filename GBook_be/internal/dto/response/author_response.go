package response

import (
	"time"
)

type AuthorResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Biography   string    `json:"biography"`
	DateOfBirth time.Time `json:"date_of_birth"`
	Nationality string    `json:"nationality"`
}
