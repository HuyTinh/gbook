package models

import "time"

// Payment đại diện cho thông tin thanh toán trong hệ thống.
type Payment struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"`                      // Khóa chính, tự động tăng
	OrderID       uint      `json:"order_id"`                                                // ID của đơn hàng
	PaymentDate   time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP" json:"payment_date"` // Ngày thanh toán
	Amount        float64   `gorm:"not null" json:"amount"`                                  // Số tiền thanh toán, không được để trống
	PaymentMethod string    `gorm:"size:50" json:"payment_method"`                           // Phương thức thanh toán
	Order         Order     `gorm:"foreignKey:OrderID" json:"-"`                             // Thông tin đơn hàng, không xuất ra JSON
}
