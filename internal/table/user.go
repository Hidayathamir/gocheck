package table

import "gorm.io/gorm"

// User -.
type User struct {
	gorm.Model
	Username string `gorm:"column:username"`
	FullName string `gorm:"column:full_name"`
	Balance  int    `gorm:"column:balance"`
}

// UserCredential -.
type UserCredential struct {
	gorm.Model
	UserID   uint   `gorm:"column:user_id"`
	Password string `gorm:"column:password"`
}
