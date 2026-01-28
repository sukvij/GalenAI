package model

import "time"

type Employee struct {
	ID        uint    `gorm:"primaryKey" json:"id"`
	FullName  string  `gorm:"column:full_name;type:varchar(255);not null" json:"full_name"`
	JobTitle  string  `gorm:"column:job_title;type:varchar(150);not null" json:"job_title"`
	Country   string  `gorm:"type:varchar(100);not null" json:"country"`
	Salary    float64 `gorm:"type:numeric(12,2);not null" json:"salary"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
