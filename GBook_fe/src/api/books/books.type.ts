import type { IAuthor } from "../authors/authors.type";
import type { IGenre } from "../genres/genres.type";


export interface IBook {
    id: number; // ID của sách (khóa chính, tự động tăng)
    name: string; // Tên của sách
    slug: string; // Đường dẫn thân thiện (slug) của sách
    price: number; // Giá của sách
    stock_quantity: number; // Số lượng sách còn trong kho
    published_date: Date; // Ngày phát hành sách
    isbn: string; // Mã số sách ISBN
    description: string; // Mô tả nội dung sách
    cover_image_url: string; // URL của ảnh bìa sách
    author: IAuthor; // Thông tin tác giả của sách
    genre: IGenre; // Thể loại của sách
}