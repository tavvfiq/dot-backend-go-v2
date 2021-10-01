package model

import "time"

type Comment struct {
	ID        int64
	UserId    int64   `gorm:"column:user_id"`
	User      User    `gorm:"foreignKey:user_id;references:id"`
	ArticleId int64   `gorm:"column:article_id"`
	Article   Article `gorm:"foreignKey:article_id;references:id;constraint:OnDelete:CASCADE"`
	Content   string
	CreatedAt time.Time  `gorm:"type:timestamp(0)"`
	DeletedAt *time.Time `gorm:"type:timestamp(0)"`
}
