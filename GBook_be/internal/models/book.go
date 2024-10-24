package models

import "time"

type Book struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `gorm:"size:255;not null"`
	AuthorID      uint   `gorm:"not null"`
	GenreID       uint
	Price         float64 `gorm:"not null"`
	StockQuantity int     `gorm:"not null"`
	PublishedDate time.Time
	ISBN          string `gorm:"size:13;unique"`
	Dessctiption  string `gorm:"type:text"`
	CoverImageURL string `gorm:"size:255"`
	Author        Author `gorm:"foreignKey:AuthorID"`
	Genre         Genre  `gorm:"foreignKey:GenreID"`
}
