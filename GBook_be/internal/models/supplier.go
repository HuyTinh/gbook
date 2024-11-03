package models

type Supplier struct {
	ID            uint           `gorm:"primaryKey;autoIncrement"`
	SupplierName  string         `gorm:"size:255;not null"`
	ContactName   string         `gorm:"size:255"`
	Address       string         `gorm:"type:text"`
	City          string         `gorm:"size:100"`
	Country       string         `gorm:"size:100"`
	PhoneNumber   string         `gorm:"size:20"`
	BookSuppliers []BookSupplier `gorm:"foreignKey:SupplierID"`
}
