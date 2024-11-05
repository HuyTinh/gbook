package models

import "time"

// Order đại diện cho đơn hàng trong hệ thống.
type Order struct {
	ID              uint          `gorm:"primaryKey;autoIncrement" json:"id"`      // Khóa chính, tự động tăng
	CustomerID      uint          `json:"customer_id"`                             // ID của khách hàng
	OrderDate       time.Time     `gorm:"type:date" json:"order_date"`             // Ngày đặt hàng
	TotalAmount     float64       `gorm:"not null" json:"total_amount"`            // Tổng số tiền, không được để trống
	Status          string        `gorm:"size:50" json:"status"`                   // Trạng thái đơn hàng
	ShippingAddress string        `gorm:"type:text" json:"shipping_address"`       // Địa chỉ giao hàng
	Customer        Customer      `gorm:"foreignKey:CustomerID" json:"-"`          // Thông tin khách hàng, không xuất ra JSON
	OrderDetails    []OrderDetail `gorm:"foreignKey:OrderID" json:"order_details"` // Danh sách chi tiết đơn hàng
}
