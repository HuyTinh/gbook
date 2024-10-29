package models

import "time"

type BookSupplier struct {
	ID          uint `gorm:"primaryKey"`
	BookID      uint
	SupplierID  uint
	SupplyPrice float64   `gorm:"not null"`
	SupplyDate  time.Time `gorm:"default:CURRENT_TIMESTAMP"`
	Book        Book      `gorm:"foreignKey:BookID"`
	Supplier    Supplier  `gorm:"foreignKey:SupplierID"`
}
