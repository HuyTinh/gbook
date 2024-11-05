package database

import (
	"GBook_be/internal/models"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	sqlite "github.com/glebarez/sqlite" // Sử dụng gói sqlite để làm việc với cơ sở dữ liệu SQLite
	"gorm.io/gorm"                      // Gói GORM để tương tác với cơ sở dữ liệu
)

// Service định nghĩa một dịch vụ tương tác với cơ sở dữ liệu.
type Service interface {
	// Health trả về một bản đồ chứa thông tin trạng thái sức khỏe.
	// Các khóa và giá trị trong bản đồ là riêng cho dịch vụ.
	Health() map[string]string

	// Close kết thúc kết nối đến cơ sở dữ liệu.
	// Nó trả về lỗi nếu không thể đóng kết nối.
	Close() error
}

var (
	dbname     = os.Getenv("GOBOOK_DB_DATABASE") // Lấy tên cơ sở dữ liệu từ biến môi trường
	dbInstance *gorm.DB                          // Biến toàn cục để lưu trữ kết nối cơ sở dữ liệu
)

// ProvideDatabase cung cấp kết nối đến cơ sở dữ liệu.
func ProvideDatabase() *gorm.DB {
	// Tái sử dụng kết nối

	var err error

	// Kiểm tra xem kết nối đã tồn tại chưa
	if dbInstance != nil {
		return dbInstance // Trả về kết nối hiện tại nếu đã tồn tại
	}

	// Mở kết nối mới đến cơ sở dữ liệu SQLite
	gormDB, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", dbname)), &gorm.Config{})

	if err != nil {
		// Đây sẽ không phải là lỗi kết nối, mà là lỗi phân tích DSN hoặc
		// một lỗi khởi tạo khác.
		log.Fatal(err)
	}

	// Tự động tạo và cập nhật bảng
	error := autoMigrate(gormDB)

	if error != nil {
		// Đây sẽ không phải là lỗi kết nối, mà là lỗi phân tích DSN hoặc
		// một lỗi khởi tạo khác.
		log.Fatal("Cannot create table ", error)
	}

	sqlDB, err := gormDB.DB()
	if err != nil {
		log.Fatal(err)
	}

	// Cấu hình kết nối
	sqlDB.SetMaxIdleConns(10)           // Số kết nối nhàn rỗi tối đa
	sqlDB.SetMaxOpenConns(100)          // Số kết nối tối đa
	sqlDB.SetConnMaxLifetime(time.Hour) // Thời gian tối đa cho một kết nối

	dbInstance = gormDB // Lưu kết nối vào biến toàn cục

	return gormDB // Trả về kết nối cơ sở dữ liệu
}

// Health kiểm tra sức khỏe của kết nối cơ sở dữ liệu bằng cách ping vào cơ sở dữ liệu.
// Nó trả về một bản đồ với các khóa cho biết các thông số sức khỏe khác nhau.
func Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	sqlDB, _ := dbInstance.DB() // Lấy đối tượng sql.DB từ GORM

	// Ping vào cơ sở dữ liệu
	err := sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// Cơ sở dữ liệu đang hoạt động, thêm nhiều thông số
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Lấy thống kê cơ sở dữ liệu (như kết nối mở, đang sử dụng, nhàn rỗi, v.v.)
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Đánh giá thống kê để cung cấp thông điệp sức khỏe
	if dbStats.OpenConnections > 40 { // Giả sử 50 là tối đa cho ví dụ này
		stats["message"] = "The database is experiencing heavy load."
	}
	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Hàm tự động gọi AutoMigrate cho tất cả các model
func autoMigrate(db *gorm.DB) error {
	// Tự động tạo bảng cho các mô hình
	error := db.AutoMigrate(
		&models.Book{},
		&models.Author{},
		&models.Customer{},
		&models.Genre{},
		&models.Order{},
		&models.OrderDetail{},
		&models.Payment{},
		&models.Review{},
	)

	return error // Trả về lỗi nếu có
}

// Close đóng kết nối cơ sở dữ liệu.
// Nó ghi lại một thông điệp chỉ ra việc ngắt kết nối khỏi cơ sở dữ liệu cụ thể.
// Nếu kết nối được đóng thành công, nó trả về nil.
// Nếu có lỗi xảy ra khi đóng kết nối, nó trả về lỗi.
func Close() error {
	sqlDB, _ := dbInstance.DB()                          // Lấy đối tượng sql.DB từ GORM
	log.Printf("Disconnected from database: %s", dbname) // Ghi lại thông điệp ngắt kết nối
	return sqlDB.Close()                                 // Đóng kết nối
}
