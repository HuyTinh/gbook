package response

type APIResponse[T any] struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
	Data    T      `json:"data,omitempty"`
}

func InitializeAPIResponse[T any](code int, message string, data T) *APIResponse[T] {
	return &APIResponse[T]{
		Code:    code,
		Message: message,
		Data:    data,
	}
}
