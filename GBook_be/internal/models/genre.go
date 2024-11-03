package models

type Genre struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	GenreName string `gorm:"size:100;not null"`
	Books     []Book `gorm:"foreignKey:GenreID"`
}
