import axiosClient from "../api-client"

export const booksService = {
    getAll: async () => {
        try {
            const response = await axiosClient.get("/books")
            return response;
        } catch (error) {
            throw error
        }
    }
}