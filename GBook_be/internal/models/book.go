package models

import "time"

type Book struct {
	ID            uint   `gorm:"primaryKey;autoIncrement"`
	Name          string `gorm:"size:255;not null"`
	Slug          string `gorm:"size:255;not null"`
	AuthorID      uint   `gorm:"not null"`
	GenreID       uint
	Price         float64   `gorm:"not null"`
	StockQuantity int       `gorm:"not null"`
	PublishedDate time.Time `gorm:"type:date"`
	ISBN          string    `gorm:"size:13;unique"`
	Description   string    `gorm:"type:text"`
	CoverImageURL string    `gorm:"size:255"`
	Author        Author    `gorm:"foreignKey:AuthorID"`
	Genre         Genre     `gorm:"foreignKey:GenreID"`
}
