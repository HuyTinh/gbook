package models

import "time"

// Review đại diện cho đánh giá sách trong hệ thống.
type Review struct {
	ID         uint      `gorm:"primaryKey;autoIncrement" json:"id"`                     // Khóa chính, tự động tăng
	BookID     uint      `json:"book_id"`                                                // ID của sách
	CustomerID uint      `json:"customer_id"`                                            // ID của khách hàng
	Rating     int       `gorm:"check:rating>=1 AND rating<=5" json:"rating"`            // Đánh giá từ 1 đến 5
	ReviewText string    `gorm:"type:text" json:"review_text"`                           // Nội dung đánh giá
	ReviewDate time.Time `gorm:"type:date;default:CURRENT_TIMESTAMP" json:"review_date"` // Ngày đánh giá
	Book       Book      `gorm:"foreignKey:BookID" json:"-"`                             // Thông tin sách, không xuất ra JSON
	Customer   Customer  `gorm:"foreignKey:CustomerID" json:"-"`                         // Thông tin khách hàng, không xuất ra JSON
}
