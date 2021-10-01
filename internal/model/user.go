package model

import "time"

type User struct {
	ID        int64
	Name      string
	CreatedAt time.Time `gorm:"type:timestamp(0)"`
	UpdatedAt time.Time `gorm:"type:timestamp(0)"`
}

func (User) TableName() string {
	return "users"
}
