// src/api/axiosInstance.js
import axios from 'axios';

const axiosClient = axios.create({
    baseURL: import.meta.env.VITE_API_URL,
});

// Interceptors để xử lý token
axiosClient.interceptors.request.use(config => {
    const token = localStorage.getItem('token'); // Lấy token từ localStorage
    if (token) {
        config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
}, error => {
    return Promise.reject(error);
});

// Interceptors để xử lý lỗi
axiosClient.interceptors.response.use(
    response => response.data,
    error => {
        // Xử lý lỗi toàn cục
        if (error.response && error.response.status === 401) {
            // Xử lý không có quyền truy cập (ví dụ: đăng xuất)
            localStorage.removeItem('token');
        }
        return Promise.reject(error);
    }
);

export default axiosClient;
