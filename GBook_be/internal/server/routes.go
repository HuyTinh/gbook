package server

import (
	"github.com/gin-contrib/cors" // Nhập gói cors để cấu hình CORS cho ứng dụng
	"github.com/gin-gonic/gin"    // Nhập gói gin để xây dựng ứng dụng web
)

// ProvideRoutes cung cấp một đối tượng gin.Engine với các route đã được cấu hình.
func ProvideRoutes() *gin.Engine {
	r := gin.Default() // Khởi tạo một đối tượng gin.Engine với cấu hình mặc định

	// Cấu hình CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},                      // Cho phép tất cả các nguồn gốc (origins)
		AllowMethods:     []string{"GET", "POST", "OPTIONS"}, // Cho phép các phương thức GET, POST và OPTIONS
		AllowHeaders:     []string{"Origin", "Content-Type"}, // Cho phép các header: Origin và Content-Type
		ExposeHeaders:    []string{"Content-Length"},         // Các header sẽ được tiết lộ cho client
		AllowCredentials: true,                               // Cho phép gửi cookie trong các yêu cầu
		MaxAge:           12 * 3600,                          // Lưu cache phản hồi preflight trong 12 giờ
	}))

	return r // Trả về đối tượng gin.Engine đã được cấu hình
}
