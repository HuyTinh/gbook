package models

import "time"

type Order struct {
	ID              uint `gorm:"primaryKey;autoIncrement"`
	CustomerID      uint
	OrderDate       time.Time     `gorm:"type:date"`
	TotalAmount     float64       `gorm:"not null"`
	Status          string        `gorm:"size:50"`
	ShippingAddress string        `gorm:"type:text"`
	Customer        Customer      `gorm:"foreignKey:CustomerID"`
	OrderDetails    []OrderDetail `gorm:"foreignKey:OrderID"`
}
