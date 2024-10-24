package models

import "time"

type Payement struct {
	ID            uint `gorm:"primaryKey"`
	OrderID       uint
	PaymentDate   time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Amount        float64   `gorm:"not null"`
	PaymentMethod string    `gorm:"size:50"`
	Order         Order     `gorm:"foreignKey:OrderID"`
}
