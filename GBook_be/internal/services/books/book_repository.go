package books

import (
	"GBook_be/internal/models" // Nhập gói mô hình để sử dụng các kiểu dữ liệu
	"errors"                   // Nhập gói errors để xử lý lỗi
	"fmt"                      // Nhập gói fmt để định dạng chuỗi
	"sync"                     // Nhập gói sync để sử dụng WaitGroup

	"gorm.io/gorm" // Nhập gói GORM để làm việc với cơ sở dữ liệu
)

// Định nghĩa interface BookRepository với các phương thức cần thiết.
type (
	BookRepository interface {
		FindAllBook() ([]models.Book, error)                 // Tìm tất cả sách
		FindBookById(bookId int64) (models.Book, error)      // Tìm sách theo ID
		FindBookBySlug(bookSlug string) (models.Book, error) // Tìm sách theo slug
		SaveBook(saveBook models.Book) (models.Book, error)  // Lưu sách mới
		// Các phương thức cập nhật và xóa sách có thể được thêm vào sau
		// UpdateBook(bookId int64, updateBook models.Book)
		// DeleteBook(bookId int64)
	}

	// Cấu trúc BookRepositoryImpl triển khai BookRepository
	BookRepositoryImpl struct {
		db *gorm.DB // Kết nối đến cơ sở dữ liệu
	}
)

// Cung cấp một thể hiện của BookRepositoryImpl
func ProvideBookRepository(db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{
		db: db, // Khởi tạo BookRepositoryImpl với kết nối cơ sở dữ liệu
	}
}

// Tìm tất cả sách và trả về danh sách sách
func (br BookRepositoryImpl) FindAllBook() ([]models.Book, error) {
	batchSize := 512 * 3 // Kích thước lô
	offset := 0          // Offset để phân trang

	var books []models.Book             // Danh sách sách
	var wg sync.WaitGroup               // Đối tượng WaitGroup để đồng bộ hóa
	results := make(chan []models.Book) // Kênh để nhận kết quả
	done := make(chan struct{})         // Kênh để báo hiệu hoàn thành

	// Goroutine để lấy sách theo từng lô
	go func() {
		for {
			batch := make([]models.Book, 0, batchSize) // Khởi tạo lô sách
			if err := br.db.Preload("Author").Preload("Genre").Limit(batchSize).Offset(offset).Find(&batch).Error; err != nil {
				close(done) // Đóng kênh nếu có lỗi
				return
			}

			if len(batch) == 0 {
				break // Thoát nếu không còn sách nào
			}

			results <- batch    // Gửi lô sách qua kênh
			offset += batchSize // Tăng offset cho lần gọi tiếp theo
		}
		close(results) // Đóng kênh kết quả khi hoàn thành
	}()

	// Goroutine để thu thập kết quả từ kênh
	go func() {
		for batch := range results {
			books = append(books, batch...) // Thêm các sách vào danh sách
		}
		close(done) // Gửi tín hiệu hoàn thành
	}()

	wg.Wait() // Đợi cho tất cả các goroutine hoàn thành
	<-done    // Chờ cho kết quả hoàn thành

	return books, nil // Trả về danh sách sách
}

// Tìm sách theo ID
func (br BookRepositoryImpl) FindBookById(bookId int64) (models.Book, error) {
	var book models.Book // Biến để lưu sách

	// Tìm sách theo ID
	if err := br.db.Preload("Author").First(&book, bookId).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Book{}, fmt.Errorf("book with ID %d not found", bookId) // Trả về lỗi nếu không tìm thấy sách
		}

		return models.Book{}, fmt.Errorf("failed to retrieve book: %v", err) // Trả về lỗi nếu có vấn đề trong quá trình truy vấn
	}

	return book, nil // Trả về sách tìm thấy
}

// Tìm sách theo slug
func (br BookRepositoryImpl) FindBookBySlug(bookSlug string) (models.Book, error) {
	var book models.Book // Biến để lưu sách

	// Tìm sách theo slug
	if err := br.db.Preload("Author").Where("slug = ?", bookSlug).First(&book).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.Book{}, fmt.Errorf("book with Slug %s not found", bookSlug) // Trả về lỗi nếu không tìm thấy sách
		}

		return models.Book{}, fmt.Errorf("failed to retrieve book: %v", err) // Trả về lỗi nếu có vấn đề trong quá trình truy vấn
	}

	return book, nil // Trả về sách tìm thấy
}

// Lưu sách mới vào cơ sở dữ liệu
func (br BookRepositoryImpl) SaveBook(saveBook models.Book) (models.Book, error) {
	var book models.Book // Biến để lưu sách

	// Lưu sách mới vào cơ sở dữ liệu
	if err := br.db.Create(&saveBook).Error; err != nil {
		return models.Book{}, err // Trả về lỗi nếu lưu không thành công
	}

	if err := br.db.First(&book, saveBook.ID).Error; err != nil {
		return models.Book{}, fmt.Errorf("create book fail") // Trả về lỗi nếu không tìm thấy sách
	}

	return book, nil // Trả về sách đã lưu
}
