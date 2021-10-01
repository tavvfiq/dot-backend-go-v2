package domain

import (
	"time"
)

type NewArticleRequestData struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	Content  string `json:"content"`
	AuthorId int64  `json:"author_id"`
}

type ArticleData struct {
	Id         int64     `json:"id"`
	Title      string    `json:"title"`
	Subtitle   string    `json:"subtitle"`
	AuthorName string    `json:"author_name"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type ArticleRequestData struct {
	Title string `json:"title"`
	Limit int    `json:"limit"`
	Page  int    `json:"page"`
}

type ArticleResponseData struct {
	Articles   []ArticleData `json:"articles"`
	Pagination Pagination    `json:"pagination"`
}

type ArticleDetailRequestData struct {
	Id int64 `json:"id"`
}

type ArticleDetailData struct {
	Id         int64            `json:"id"`
	Title      string           `json:"title"`
	Subtitle   string           `json:"subtitle"`
	AuthorId   int64            `json:"author_id"`
	AuthorName string           `json:"author_name"`
	Content    string           `json:"content"`
	Comments   []ArticleComment `json:"comments"`
}

type ArticleComment struct {
	Id        int64
	UserName  string    `json:"user_name"`
	UserId    int64     `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type UpdateArticleRequestData struct {
	Title    *string `json:"title,omitempty"`
	Subtitle *string `json:"subtitle,omitempty"`
	Content  *string `json:"content,omitempty"`
	AuthorId *int64  `json:"author_id,omitempty"`
}
