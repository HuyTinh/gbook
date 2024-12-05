import type { APIResponse } from "../../@types/api-response";
import axiosClient from "../axios-client"
import type { IBook } from "./book.type";

export const bookService = {
    getAll: async (): Promise<APIResponse<IBook[]>> => {
        try {
            const response = axiosClient.get("/books")
            return response as unknown as Promise<APIResponse<IBook[]>>
        } catch (error) {
            console.log(error);
            return [] as unknown as Promise<APIResponse<IBook[]>>
        }
    }
}