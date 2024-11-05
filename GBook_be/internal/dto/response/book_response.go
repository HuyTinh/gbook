package response

import (
	"GBook_be/internal/models"
	"time"
)

// BookResponse đại diện cho thông tin phản hồi về sách trong hệ thống.
type BookResponse struct {
	ID            uint          `json:"id"`              // ID của sách (khóa chính, tự động tăng)
	Name          string        `json:"name"`            // Tên của sách
	Slug          string        `json:"slug"`            // Đường dẫn thân thiện (slug) của sách
	Price         float64       `json:"price"`           // Giá của sách
	StockQuantity int           `json:"stock_quantity"`  // Số lượng sách còn trong kho
	PublishedDate time.Time     `json:"published_date"`  // Ngày phát hành sách
	ISBN          string        `json:"isbn"`            // Mã số sách ISBN
	Description   string        `json:"description"`     // Mô tả nội dung sách
	CoverImageURL string        `json:"cover_image_url"` // URL của ảnh bìa sách
	Author        models.Author `json:"author"`          // Thông tin tác giả của sách
	Genre         models.Genre  `json:"genre"`           // Thể loại của sách
}
