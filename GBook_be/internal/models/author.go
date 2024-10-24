package models

import "time"

type Author struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"size:255;not null"`
	Biography   string `gorm:"type:text"`
	DateOfBirth time.Time
	Nationality string `gorm:"size:100"`
	Books       []Book `gorm:"foreignKey:AuthorID"`
}
