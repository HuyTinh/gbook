package models

// Customer đại diện cho khách hàng trong hệ thống.
type Customer struct {
	ID          uint    `gorm:"primaryKey;autoIncrement" json:"id"`  // Khóa chính, tự động tăng
	FirstName   string  `gorm:"size:100;not null" json:"first_name"` // Tên, không được để trống
	LastName    string  `gorm:"size:100;not null" json:"last_name"`  // Họ, không được để trống
	Email       string  `gorm:"size:255;unique" json:"email"`        // Địa chỉ email, duy nhất
	PhoneNumber string  `gorm:"size:20" json:"phone_number"`         // Số điện thoại
	Address     string  `gorm:"type:text" json:"address"`            // Địa chỉ chi tiết
	City        string  `gorm:"size:100" json:"city"`                // Thành phố
	Country     string  `gorm:"size:100" json:"country"`             // Quốc gia
	Orders      []Order `gorm:"foreignKey:CustomerID" json:"orders"` // Danh sách đơn hàng của khách hàng
}
