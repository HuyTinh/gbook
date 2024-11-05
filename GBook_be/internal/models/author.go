package models

import "time"

// Author đại diện cho tác giả trong hệ thống.
type Author struct {
	ID          uint      `gorm:"primaryKey;autoIncrement" json:"id"` // Khóa chính, tự động tăng
	Name        string    `gorm:"size:255;not null" json:"name"`      // Tên tác giả, không được để trống
	Biography   string    `gorm:"type:text" json:"biography"`         // Tiểu sử của tác giả
	DateOfBirth time.Time `gorm:"type:datetime" json:"date_of_birth"` // Ngày sinh của tác giả
	Nationality string    `gorm:"size:100" json:"nationality"`        // Quốc tịch của tác giả
	Books       []Book    `gorm:"foreignKey:AuthorID" json:"-"`       // Danh sách sách của tác giả, không trả về trong JSON
}
