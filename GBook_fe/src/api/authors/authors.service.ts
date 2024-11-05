import axiosClient from "../api-client"

export const authorsService = {
    getAll: async () => {
        try {
            const response = await axiosClient.get("/authors")
            return response;
        } catch (error) {
            throw error
        }
    }
}