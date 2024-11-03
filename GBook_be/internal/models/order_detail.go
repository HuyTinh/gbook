package models

type OrderDetail struct {
	ID       uint `gorm:"primaryKey;autoIncrement"`
	OrderID  uint
	BookID   uint
	Quantity int     `gorm:"not null"`
	UnitPage float64 `gorm:"not null"`
	Order    Order   `gorm:"foreignKey:OrderID"`
	Book     Book    `gorm:"foreignKey:BookID"`
}
