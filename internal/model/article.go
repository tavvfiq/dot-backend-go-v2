package model

import "time"

type Article struct {
	ID        int64
	Title     string
	Subtitle  string
	AuthorId  int64 `gorm:"column:author_id"`
	Author    User  `gorm:"foreignKey:author_id;references:id;constraint:OnDelete:CASCADE"`
	Content   string
	CreatedAt time.Time  `gorm:"type:timestamp(0)"`
	UpdatedAt time.Time  `gorm:"type:timestamp(0)"`
	DeletedAt *time.Time `gorm:"type:timestamp(0)"`
}
