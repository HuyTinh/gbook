package response

// APIResponse đại diện cho phản hồi API với cấu trúc chung.
// T là kiểu dữ liệu tổng quát, cho phép chứa bất kỳ loại dữ liệu nào.
type APIResponse[T any] struct {
	Code    int    `json:"code"`                // Mã trạng thái HTTP hoặc mã tùy chọn của API.
	Message string `json:"message,omitempty"`   // Thông điệp mô tả phản hồi, có thể để trống.
	Data    T      `json:"data,omitempty"`      // Dữ liệu được trả về, có thể để trống.
}

// InitializeAPIResponse khởi tạo một phản hồi API với mã, thông điệp và dữ liệu cho trước.
func InitializeAPIResponse[T any](code int, message string, data T) *APIResponse[T] {
	return &APIResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
