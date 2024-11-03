package response

import (
	"GBook_be/internal/models"
	"time"
)

type BookResponse struct {
	ID            uint          `json:"id"`
	Name          string        `json:"name"`
	Slug          string        `json:"slug"`
	Price         float64       `json:"price"`
	StockQuantity int           `json:"stock_quantity"`
	PublishedDate time.Time `gorm:"type:date"`     `json:"published_date"`
	ISBN          string        `json:"ibsn"`
	Description   string        `json:"description"`
	CoverImageURL string        `json:"cover_image_URL"`
	Author        models.Author `json:"author"`
	Genre         models.Genre  `json:"genre"`
}
