package domain

type AddCommentRequest struct {
	UserId    int64  `json:"user_id,omitempty"`
	ArticleId int64  `json:"article_id"`
	Content   string `json:"content"`
}
