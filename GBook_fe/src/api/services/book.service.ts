import type { IBook } from "../../types/book.type"
import api from "../axios-client"

export const BookService = {
    getAllBook: async (): Promise<IBook[]> => {
        try {
            return await api.get("/books")
        } catch (err) {
            return [] as IBook[]
        }
    },
    getBookById: async (bookId: number): Promise<IBook> => {
        try {
            return await api.get(`/books/${bookId}`) as IBook
        } catch (err) {
            return {} as IBook
        }
    },
    createBook: async (newBook: { title: string, genre: string, authorId: number }): Promise<IBook> => {
        try {
            return await api.post(`/books`, newBook) as IBook
        } catch (err) {
            return {} as IBook
        }
    }
}