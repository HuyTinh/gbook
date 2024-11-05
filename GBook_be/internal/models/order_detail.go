package models

// OrderDetail đại diện cho chi tiết đơn hàng trong hệ thống.
type OrderDetail struct {
	ID       uint    `gorm:"primaryKey;autoIncrement" json:"id"` // Khóa chính, tự động tăng
	OrderID  uint    `json:"order_id"`                           // ID của đơn hàng
	BookID   uint    `json:"book_id"`                            // ID của sách
	Quantity int     `gorm:"not null" json:"quantity"`           // Số lượng sách, không được để trống
	UnitPage float64 `gorm:"not null" json:"unit_page"`          // Giá của mỗi cuốn sách
	Order    Order   `gorm:"foreignKey:OrderID" json:"-"`        // Thông tin đơn hàng, không xuất ra JSON
	Book     Book    `gorm:"foreignKey:BookID" json:"-"`         // Thông tin sách, không xuất ra JSON
}
