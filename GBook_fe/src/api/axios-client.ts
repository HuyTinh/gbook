import axios from 'axios';

// Create an Axios instance with default settings
const api = axios.create({
    baseURL: import.meta.env.VITE_END_POINT_URL as string, // Replace with your base URL
    timeout: 10000, // Optional: Set a timeout for requests
    headers: {
        'Content-Type': 'application/json',
    },
});

// Optionally, you can add interceptors if needed
api.interceptors.response.use(
    response => response.data.data,
    error => {
        // Handle errors globally
        console.error('API error:', error);
        return Promise.reject(error);
    }
);

// Export the Axios instance
export default api;
