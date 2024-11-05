package models

type Genre struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	GenreName string `gorm:"size:100;not null" json:"genre_name"`
	Books     []Book `gorm:"foreignKey:GenreID" json:"-"`
}
