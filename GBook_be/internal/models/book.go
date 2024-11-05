package models

import "time"

// Book đại diện cho sách trong hệ thống.
type Book struct {
	ID            uint      `gorm:"primaryKey;autoIncrement" json:"id"` // Khóa chính, tự động tăng
	Name          string    `gorm:"size:255;not null" json:"name"`      // Tên sách, không được để trống
	Slug          string    `gorm:"size:255;not null" json:"slug"`      // Đường dẫn thân thiện với URL
	AuthorID      uint      `gorm:"not null" json:"author_id"`          // ID của tác giả, không được để trống
	GenreID       uint      `json:"genre_id"`                           // ID thể loại
	Price         float64   `gorm:"not null" json:"price"`              // Giá sách, không được để trống
	StockQuantity int       `gorm:"not null" json:"stock_quantity"`     // Số lượng còn lại trong kho, không được để trống
	PublishedDate time.Time `gorm:"type:date" json:"published_date"`    // Ngày xuất bản
	ISBN          string    `gorm:"size:13;unique" json:"isbn"`         // Mã ISBN, duy nhất cho mỗi cuốn sách
	Description   string    `gorm:"type:text" json:"description"`       // Mô tả sách
	CoverImageURL string    `gorm:"size:255" json:"cover_image_url"`    // URL hình ảnh bìa sách
	Author        Author    `gorm:"foreignKey:AuthorID" json:"-"`       // Thông tin tác giả, không xuất ra JSON
	Genre         Genre     `gorm:"foreignKey:GenreID" json:"-"`        // Thông tin thể loại, không xuất ra JSON
}
