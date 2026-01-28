package model

import "time"

type Employee struct {
	ID        uint    `gorm:"primaryKey"`
	FullName  string  `gorm:"column:full_name;type:varchar(255);not null"`
	JobTitle  string  `gorm:"column:job_title;type:varchar(150);not null"`
	Country   string  `gorm:"type:varchar(100);not null"`
	Salary    float64 `gorm:"type:numeric(12,2);not null"` // gross salary
	CreatedAt time.Time
	UpdatedAt time.Time
}
