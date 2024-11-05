package main

import (
	"GBook_be/internal/server"           // Nhập gói server để cung cấp máy chủ HTTP
	"GBook_be/internal/services/authors" // Nhập gói dịch vụ cho tác giả
	"GBook_be/internal/services/books"   // Nhập gói dịch vụ cho sách
	"context"                            // Nhập gói context để quản lý ngữ cảnh cho các goroutine
	"fmt"                                // Nhập gói fmt để định dạng chuỗi
	"log"                                // Nhập gói log để ghi log
	"net/http"                           // Nhập gói http để làm việc với máy chủ HTTP
	"os/signal"                          // Nhập gói signal để nhận tín hiệu từ hệ thống
	"syscall"                            // Nhập gói syscall để làm việc với hệ thống
	"time"                               // Nhập gói time để xử lý thời gian

	"go.uber.org/fx" // Nhập gói fx để sử dụng dependency injection
)

// Hàm gracefulShutdown để tắt máy chủ một cách nhẹ nhàng
func gracefulShutdown(apiServer *http.Server, done chan bool) {
	// Tạo ngữ cảnh lắng nghe tín hiệu ngắt từ hệ điều hành
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Lắng nghe tín hiệu ngắt
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// Ngữ cảnh được sử dụng để thông báo cho máy chủ có 5 giây để hoàn thành
	// yêu cầu mà nó đang xử lý
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Thông báo cho goroutine chính rằng việc tắt máy chủ đã hoàn tất
	done <- true
}

// Hàm startServer để khởi động máy chủ HTTP
func startServer(server *http.Server) {
	// Tạo một kênh done để tín hiệu khi tắt máy chủ hoàn tất
	done := make(chan bool, 1)

	// Chạy việc tắt máy chủ một cách nhẹ nhàng trong goroutine riêng
	go gracefulShutdown(server, done)

	// Bắt đầu lắng nghe và phục vụ
	err := server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// Chờ cho việc tắt máy chủ hoàn tất
	<-done
	log.Println("Graceful shutdown complete.")
}

// Hàm main là điểm khởi đầu của ứng dụng
func main() {
	fx.New(
		server.ProvideServers(),  // Cung cấp máy chủ
		books.ProvideBooks(),     // Cung cấp dịch vụ sách
		authors.ProvideAuthors(), // Cung cấp dịch vụ tác giả
		fx.Invoke(
			server.ProvideController, // Cung cấp controller cho máy chủ
			startServer,              // Bắt đầu máy chủ
		),
	).Run() // Chạy ứng dụng với fx
}
