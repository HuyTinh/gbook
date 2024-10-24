package models

type Customer struct {
	ID          uint    `gorm:"primaryKey"`
	FirstName   string  `gorm:"size:100;not null"`
	LastName    string  `gorm:"size:100;not null"`
	Email       string  `gorm:"size:255;unique"`
	PhoneNumber string  `gorm:"size:20"`
	Address     string  `gorm:"type:text"`
	City        string  `gorm:"size:100"`
	Country     string  `gorm:"size:100"`
	Orders      []Order `gorm:"foreignKey:CustomerID"`
}
