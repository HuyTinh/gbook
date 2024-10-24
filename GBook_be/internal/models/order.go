package models

import "time"

type Order struct {
	ID              uint `gorm:"primaryKey"`
	CustomerID      uint
	OrderDate       time.Time
	TotalAmount     float64       `gorm:"not null"`
	Status          string        `gorm:"size:50"`
	ShippingAddress string        `gorm:"type:text"`
	Customer        Customer      `gorm:"foreignKey:CustomerID"`
	OrderDetails    []OrderDetail `gorm:"foreignKey:OrderID"`
}
