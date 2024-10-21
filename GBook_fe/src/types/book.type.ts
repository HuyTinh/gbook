import type { IAuthor } from "./author.type";

export interface IBook {
    id: number,
    title: string,
    genre: string,
    author: IAuthor
}