package books

import (
	APIResponse "GBook_be/internal/dto/response" // Nhập gói để xử lý phản hồi API
	"GBook_be/internal/models"                   // Nhập gói mô hình để sử dụng các kiểu dữ liệu

	"github.com/gin-gonic/gin" // Nhập gói gin để xây dựng ứng dụng web
)

// BookService chứa repository để thao tác với dữ liệu sách.
type BookService struct {
	repository BookRepository // Repository cho sách
}

// ProvideBookService cung cấp BookService với repository đã cho.
func ProvideBookService(bookRepository BookRepository) BookService {
	return BookService{
		repository: bookRepository, // Khởi tạo BookService với repository
	}
}

// GetAllBook lấy tất cả sách và trả về dưới dạng JSON.
func (bs *BookService) GetAllBook(c *gin.Context) {
	// Tìm tất cả sách từ repository
	result, err := bs.repository.FindAllBook()

	// Kiểm tra lỗi khi lấy sách
	if err != nil {
		c.JSON(500, APIResponse.InitializeAPIResponse(500, "Lỗi khi lấy sách: "+err.Error(), ""))
		return
	}

	// Kiểm tra nếu không có sách nào
	if len(result) == 0 {
		c.JSON(200, APIResponse.InitializeAPIResponse(200, "Không tìm thấy sách nào", ""))
		return
	}

	bookResponsesChan := make(chan APIResponse.BookResponse) // Kênh để truyền thông tin phản hồi sách
	done := make(chan bool)                                  // Kênh để báo hiệu khi hoàn thành

	// Goroutine để chuyển đổi mô hình sách sang phản hồi API
	go func() {
		for _, book := range result {
			bookResponse := APIResponse.BookResponse{
				ID:            book.ID,
				Name:          book.Name,
				Slug:          book.Slug,
				Price:         book.Price,
				StockQuantity: book.StockQuantity,
				PublishedDate: book.PublishedDate,
				ISBN:          book.ISBN,
				Description:   book.Description,
				CoverImageURL: book.CoverImageURL,
				Author:        book.Author,
				Genre:         book.Genre,
			}

			bookResponsesChan <- bookResponse // Gửi phản hồi sách qua kênh
		}
		close(bookResponsesChan) // Đóng kênh khi hoàn thành
	}()

	bookResponses := make([]APIResponse.BookResponse, 0, len(result)) // Khởi tạo mảng để chứa phản hồi sách

	// Goroutine để thu thập phản hồi sách từ kênh
	go func() {
		for bookResponse := range bookResponsesChan {
			bookResponses = append(bookResponses, bookResponse) // Thêm phản hồi sách vào mảng
		}
		done <- true // Gửi tín hiệu hoàn thành
	}()

	<-done // Chờ cho đến khi goroutine hoàn thành

	c.JSON(200, APIResponse.InitializeAPIResponse(200, "", bookResponses)) // Trả về phản hồi API với danh sách sách
}

// SaveBook lưu một sách mới từ dữ liệu JSON.
func (bs *BookService) SaveBook(c *gin.Context) {
	var saveBook models.Book // Khai báo biến để lưu sách mới

	// Liên kết dữ liệu JSON từ yêu cầu với biến saveBook
	if err := c.ShouldBindJSON(&saveBook); err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, err.Error(), ""))
		return
	}

	// Lưu sách qua repository
	successSaveBook, err := bs.repository.SaveBook(saveBook)

	// Kiểm tra lỗi khi lưu sách
	if err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, err.Error(), ""))
		return
	}

	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", successSaveBook)) // Trả về phản hồi API thành công
}
