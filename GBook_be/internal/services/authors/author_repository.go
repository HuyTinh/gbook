package authors

import (
	"GBook_be/internal/models" // Nhập gói mô hình để sử dụng các kiểu dữ liệu
	"sync"                     // Nhập gói sync để sử dụng WaitGroup

	"gorm.io/gorm" // Nhập gói GORM để làm việc với cơ sở dữ liệu
)

// Định nghĩa interface AuthorRepository với các phương thức cần thiết.
type (
	AuthorRepository interface {
		FindAllAuthor() ([]models.Author, error)                   // Tìm tất cả tác giả
		FindAuthorById(authorId int64) (models.Author, error)      // Tìm tác giả theo ID
		FindAuthorByName(authorName string) (models.Author, error) // Tìm tác giả theo tên
		SaveAuthor(newAuthor models.Author) (models.Author, error) // Lưu tác giả mới
		// UpdateAuthor(authorId int64, updateAuthor models.Author) // Cập nhật thông tin tác giả
		// DeleteAuthor(authorId int64) // Xóa tác giả theo ID
	}

	// Cấu trúc AuthorRepositoryImpl triển khai AuthorRepository
	AuthorRepositoryImpl struct {
		db *gorm.DB // Kết nối đến cơ sở dữ liệu
	}
)

// Cung cấp một thể hiện của AuthorRepositoryImpl
func ProvideAuthorRepository(db *gorm.DB) AuthorRepository {
	return &AuthorRepositoryImpl{
		db: db, // Khởi tạo AuthorRepositoryImpl với kết nối cơ sở dữ liệu
	}
}

// Tìm tất cả tác giả và trả về danh sách tác giả
func (ar AuthorRepositoryImpl) FindAllAuthor() ([]models.Author, error) {
	batchSize := 512 * 3 // Kích thước lô

	var authors []models.Author           // Danh sách tác giả
	offset := 0                           // Offset để phân trang
	var wg sync.WaitGroup                 // Đối tượng WaitGroup để đồng bộ hóa
	results := make(chan []models.Author) // Kênh để nhận kết quả
	done := make(chan struct{})           // Kênh để báo hiệu hoàn thành

	// Goroutine để lấy tác giả theo từng lô
	go func() {
		for {
			batch := make([]models.Author, 0, batchSize) // Khởi tạo lô tác giả
			if err := ar.db.Preload("Books").Limit(batchSize).Offset(offset).Find(&batch).Error; err != nil {
				close(done) // Đóng kênh nếu có lỗi
				return
			}

			if len(batch) == 0 {
				break // Thoát nếu không còn tác giả nào
			}

			results <- batch    // Gửi lô tác giả qua kênh
			offset += batchSize // Tăng offset cho lần gọi tiếp theo
		}
		close(results) // Đóng kênh kết quả khi hoàn thành
	}()

	// Goroutine để thu thập kết quả từ kênh
	go func() {
		for batch := range results {
			authors = append(authors, batch...) // Thêm các tác giả vào danh sách
		}
		close(done) // Gửi tín hiệu hoàn thành
	}()

	// Chờ cho tất cả các lô được xử lý
	wg.Wait() // Đợi cho tất cả goroutine hoàn thành
	<-done    // Chờ cho kết quả hoàn thành

	return authors, nil // Trả về danh sách tác giả
}

// Tìm tác giả theo ID
func (ar AuthorRepositoryImpl) FindAuthorById(authorId int64) (models.Author, error) {
	var author models.Author // Biến để lưu tác giả

	// Tìm tác giả theo ID
	if err := ar.db.Preload("Books").Find(&author, authorId).Error; err != nil {
		return models.Author{}, err // Trả về lỗi nếu không tìm thấy
	}

	return author, nil // Trả về tác giả tìm thấy
}

// Tìm tác giả theo tên
func (ar AuthorRepositoryImpl) FindAuthorByName(authorName string) (models.Author, error) {
	var author models.Author // Biến để lưu tác giả

	// Tìm tác giả theo tên
	if err := ar.db.Preload("Books").Where("name = ?", authorName).First(&author).Error; err != nil {
		return models.Author{}, err // Trả về lỗi nếu không tìm thấy
	}

	return author, nil // Trả về tác giả tìm thấy
}

// Lưu tác giả mới vào cơ sở dữ liệu
func (ar AuthorRepositoryImpl) SaveAuthor(saveAuthor models.Author) (models.Author, error) {
	var author models.Author // Biến để lưu tác giả

	// Kiểm tra xem tác giả đã tồn tại chưa
	if err := ar.db.Find(&author, saveAuthor.ID).Error; err != nil {
		return models.Author{}, err // Trả về lỗi nếu không tìm thấy
	}

	// Lưu tác giả mới vào cơ sở dữ liệu
	if err := ar.db.Create(&saveAuthor).Error; err != nil {
		return models.Author{}, err // Trả về lỗi nếu lưu không thành công
	}

	return author, nil // Trả về tác giả đã lưu
}
