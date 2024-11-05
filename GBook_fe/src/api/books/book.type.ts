import type { IAuthor } from "../authors/author.type";
import type { IGenre } from "../genres/genre.type";

export interface IBook {
    id: number;
    name: string;
    slug: string;
    price: number;
    stock_quantity: number;
    published_date: string;
    isbn: string;
    description: string;
    cover_image_url: string;
    author: IAuthor;
    genre: IGenre;

}