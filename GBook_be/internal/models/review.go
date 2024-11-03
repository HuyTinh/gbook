package models

import "time"

type Review struct {
	ID         uint `gorm:"primaryKey;autoIncrement"`
	BookID     uint
	CustomerID uint
	Rating     int       `gorm:"check:rating>=1 AND rating<=5"`
	ReviewText string    `gorm:"type:text"`
	ReviewDate time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP"`
	Book       Book      `gorm:"foreignKey:BookID"`
	Customer   Customer  `gorm:"foreignKey:CustomerID"`
}
