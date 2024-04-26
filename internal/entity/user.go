package entity

import (
	"time"
)

// User -.
type User struct {
	ID uint `gorm:"primarykey" json:"id"`

	Username string `gorm:"column:username"  json:"username"`
	FullName string `gorm:"column:full_name" json:"full_name"`
	Balance  int    `gorm:"column:balance"   json:"balance"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

// UserCredential -.
type UserCredential struct {
	ID uint `gorm:"primarykey" json:"id"`

	UserID   uint   `gorm:"column:user_id"  json:"user_id"`
	Password string `gorm:"column:password" json:"password"`

	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
