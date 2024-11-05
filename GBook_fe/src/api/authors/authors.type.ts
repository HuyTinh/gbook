// AuthorResponse đại diện cho thông tin phản hồi về tác giả trong hệ thống.
export interface IAuthor {
    id: number;            // ID của tác giả (khóa chính, tự động tăng)
    name: string;         // Tên của tác giả
    biography: string;    // Tiểu sử của tác giả
    date_of_birth: Date;    // Ngày sinh của tác giả
    nationality: string;   // Quốc tịch của tác giả
}

