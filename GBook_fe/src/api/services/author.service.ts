import type { IAuthor } from "../../types/author.type";
import api from "../axios-client";

export const AuthorService = {
    getAllAuthor: async (): Promise<IAuthor[]> => {
        try {
            return await api.get("/authors") as IAuthor[]
        } catch (error) {
            console.log(error);
            return [] as IAuthor[]
        }
    }
}