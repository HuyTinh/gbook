package authors

import (
	APIResponse "GBook_be/internal/dto/response"
	"GBook_be/internal/models"

	"github.com/gin-gonic/gin"
)

type AuthorService struct {
	repository AuthorRepository
}

func ProvideAuthorService(bookRepository AuthorRepository) AuthorService {
	return AuthorService{
		repository: bookRepository,
	}
}

func (bc *AuthorService) GetAllAuthor(c *gin.Context) {
	// Gọi tìm tất cả tác giả với các trường cần thiết
	result, err := bc.repository.FindAllAuthor() // Giả sử bạn đã tạo phương thức này
	if err != nil {
		c.JSON(500, APIResponse.InitializeAPIResponse(500, "Error fetching authors: "+err.Error(), ""))
		return
	}

	// Kiểm tra nếu không có tác giả nào
	if len(result) == 0 {
		c.JSON(200, APIResponse.InitializeAPIResponse(200, "No authors found", ""))
		return
	}

	// Khởi tạo channel để nhận kết quả
	authorResponsesChan := make(chan APIResponse.AuthorResponse)
	done := make(chan bool)

	// Goroutine để xử lý từng tác giả
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
			authorResponsesChan <- authorResponse // Gửi kết quả vào channel
		}
		close(authorResponsesChan) // Đóng channel sau khi gửi hết kết quả
	}()

	// Khởi tạo slice để lưu trữ kết quả
	authorResponses := make([]APIResponse.AuthorResponse, 0, len(result))

	// Nhận kết quả từ channel
	go func() {
		for authorResponse := range authorResponsesChan {
			authorResponses = append(authorResponses, authorResponse) // Nhận và thêm vào slice
		}
		done <- true // Báo hiệu đã hoàn thành
	}()

	// Chờ cho đến khi tất cả đã hoàn thành
	<-done

	// Trả về kết quả
	c.JSON(200, APIResponse.InitializeAPIResponse(200, "", authorResponses))
}

func (as *AuthorService) SaveAuthor(c *gin.Context) {

	var saveAuthor models.Author

	if err := c.ShouldBindJSON(&saveAuthor); err != nil {
		c.JSON(400, APIResponse.InitializeAPIResponse(400, "Invalid input", ""))
		return
	}
	as.repository.SaveAuthor(saveAuthor)
	c.JSON(200, APIResponse.InitializeAPIResponse(1000, "", ""))
}
