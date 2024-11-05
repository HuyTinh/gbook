package authors

import (
	APIResponse "GBook_be/internal/dto/response" // Nhập gói để xử lý phản hồi API
	"GBook_be/internal/models"                   // Nhập gói mô hình để sử dụng các kiểu dữ liệu

	"github.com/gin-gonic/gin" // Nhập gói gin để xây dựng ứng dụng web
)

// AuthorService chứa repository để thao tác với dữ liệu tác giả.
type AuthorService struct {
	repository AuthorRepository // Repository cho tác giả
}

// ProvideAuthorService cung cấp AuthorService với repository đã cho.
func ProvideAuthorService(bookRepository AuthorRepository) AuthorService {
	return AuthorService{
		repository: bookRepository, // Khởi tạo AuthorService với repository
	}
}

// GetAllAuthor lấy tất cả tác giả và trả về dưới dạng JSON.
func (bc *AuthorService) GetAllAuthor(c *gin.Context) {
	// Tìm tất cả tác giả từ repository
	result, err := bc.repository.FindAllAuthor()

	// Kiểm tra lỗi khi lấy tác giả
	if err != nil {
		c.JSON(500, APIResponse.InitializeAPIResponse(500, "Lỗi khi lấy tác giả: "+err.Error(), ""))
		return
	}

	// Kiểm tra nếu không có tác giả nào
	if len(result) == 0 {
		c.JSON(200, APIResponse.InitializeAPIResponse(200, "Không tìm thấy tác giả nào", ""))
		return
	}

	authorResponsesChan := make(chan APIResponse.AuthorResponse) // Kênh để truyền thông tin phản hồi tác giả
	done := make(chan bool)                                      // Kênh để báo hiệu khi hoàn thành

	// Goroutine để chuyển đổi mô hình tác giả sang phản hồi API
	go func() {
		for _, author := range result {
			authorResponse := APIResponse.AuthorResponse{
				ID:          author.ID,
				Name:        author.Name,
				Biography:   author.Biography,
				DateOfBirth: author.DateOfBirth,
				Nationality: author.Nationality,
				// Thêm các trường khác nếu cần
			}
			authorResponsesChan <- authorResponse // Gửi phản hồi tác giả qua kênh
		}
		close(authorResponsesChan) // Đóng kênh khi hoàn thành
	}()

	// Khởi tạo slice để lưu trữ kết quả
	authorResponses := make([]APIResponse.AuthorResponse, 0, len(result))

	// Goroutine để thu thập phản hồi tác giả từ kênh
	go func() {
		for authorResponse := range authorResponsesChan {
			authorResponses = append(authorResponses, authorResponse) // Thêm phản hồi tác giả vào mảng
		}
		done <- true // Gửi tín hiệu hoàn thành
	}()

	// Chờ cho đến khi goroutine hoàn thành
	<-done

	// Trả về phản hồi API với danh sách tác giả
	c.JSON(200, APIResponse.InitializeAPIResponse(200, "", authorResponses))
}

// SaveAuthor lưu một tác giả mới từ dữ liệu JSON.
func (as *AuthorService) SaveAuthor(c *gin.Context) {
	var saveAuthor models.Author // Khai báo biến để lưu tác giả mới

	// Liên kết dữ liệu JSON từ yêu cầu với biến saveAuthor
	if err := c.ShouldBindJSON(&saveAuthor); err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, "Dữ liệu đầu vào không hợp lệ", ""))
		return
	}

	// Lưu tác giả qua repository
	as.repository.SaveAuthor(saveAuthor)

	// Trả về phản hồi API thành công
	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", ""))
}
