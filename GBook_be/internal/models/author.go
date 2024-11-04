package models

import "time"

type Author struct {
	ID          uint      `gorm:"primaryKey;autoIncrement"`
	Name        string    `gorm:"size:255;not null"`
	Biography   string    `gorm:"type:text"`
	DateOfBirth time.Time `gorm:"type:datetime"`
	Nationality string    `gorm:"size:100"`
	Books       []Book    `gorm:"foreignKey:AuthorID" json:"-"`
}
