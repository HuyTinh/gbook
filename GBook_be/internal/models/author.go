package models

type Author struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	Name      string `json:"name"`
	BirthYear string `json:"birth_year"`
}
