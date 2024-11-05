package server

import (
	"fmt"
	"net/http" // Nhập gói http để sử dụng các đối tượng và phương thức liên quan đến server HTTP
	"os"       // Nhập gói os để truy cập các biến môi trường
	"strconv"  // Nhập gói strconv để chuyển đổi giữa các kiểu dữ liệu
	"time"     // Nhập gói time để quản lý thời gian

	"github.com/gin-gonic/gin"            // Nhập gói gin để xây dựng ứng dụng web
	_ "github.com/joho/godotenv/autoload" // Tự động tải các biến môi trường từ file .env
)

// Server đại diện cho server HTTP.
type Server struct {
	port int // Cổng mà server sẽ lắng nghe
}

// ProvideServer cung cấp một server HTTP với cấu hình đã chỉ định.
func ProvideServer(r *gin.Engine) *http.Server {
	// Lấy giá trị cổng từ biến môi trường "PORT" và chuyển đổi nó sang kiểu int
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	NewServer := &Server{
		port: port, // Khởi tạo đối tượng Server với cổng đã lấy
	}

	// Định nghĩa cấu hình cho server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port), // Địa chỉ mà server sẽ lắng nghe
		Handler:      r,                                  // Bộ xử lý cho các yêu cầu HTTP
		IdleTimeout:  time.Minute,                        // Thời gian chờ khi không có yêu cầu
		ReadTimeout:  10 * time.Second,                   // Thời gian chờ khi đọc yêu cầu
		WriteTimeout: 30 * time.Second,                   // Thời gian chờ khi ghi phản hồi
	}

	return server // Trả về đối tượng server đã cấu hình
}
