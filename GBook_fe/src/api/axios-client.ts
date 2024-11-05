import axios from 'axios';

// Tạo instance Axios với cấu hình mặc định
const axiosClient = axios.create({
    baseURL: import.meta.env.VITE_API_PREFIX, // Thay đổi thành URL của API của bạn
    timeout: 10000, // Thời gian chờ (ms)
    headers: {
        'Content-Type': 'application/json',
        // Thêm các header tùy chỉnh nếu cần
    },
});

// Tùy chỉnh response interceptor nếu cần
axiosClient.interceptors.response.use(
    response => response.data,
    error => {
        // Xử lý lỗi ở đây
        console.error('API error:', error);
        return Promise.reject(error);
    }
);

// Xuất instance Axios
export default axiosClient;
