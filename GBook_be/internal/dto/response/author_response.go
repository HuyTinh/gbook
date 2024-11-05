package response

import (
	"time"
)

// AuthorResponse đại diện cho thông tin phản hồi về tác giả trong hệ thống.
type AuthorResponse struct {
	ID          uint      `json:"id"`            // ID của tác giả (khóa chính, tự động tăng)
	Name        string    `json:"name"`          // Tên của tác giả
	Biography   string    `json:"biography"`     // Tiểu sử của tác giả
	DateOfBirth time.Time `json:"date_of_birth"` // Ngày sinh của tác giả
	Nationality string    `json:"nationality"`   // Quốc tịch của tác giả
}
