package user

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserName  string `gorm:"uniqueIndex;not null" json:"user_name"`
	Password  string `gorm:"not null" json:"password"`
	Role      string `json:"role"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Credentials struct {
	UserName string `gorm:"uniqueIndex;not null" json:"user_name"`
	Password string `gorm:"not null" json:"password"`
}
