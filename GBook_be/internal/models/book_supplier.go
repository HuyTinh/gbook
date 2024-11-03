package models

import "time"

type BookSupplier struct {
	ID          uint `gorm:"primaryKey;autoIncrement"`
	BookID      uint
	SupplierID  uint
	SupplyPrice float64   `gorm:"not null"`
	SupplyDate  time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP"`
	Book        Book      `gorm:"foreignKey:BookID"`
	Supplier    Supplier  `gorm:"foreignKey:SupplierID"`
}
