package model

type Salary struct {
	ID              uint    `gorm:"primaryKey" json:"id"`
	FullName        string  `gorm:"column:full_name;type:varchar(255);not null" json:"full_name"`
	JobTitle        string  `gorm:"column:job_title;type:varchar(150);not null" json:"job_title"`
	Country         string  `gorm:"type:varchar(100);not null" json:"country"`
	GrossSalary     float64 `json:"gross_salary"`
	DeductionAmount float64 `json:"deduction_amount"`
	NetSalary       float64 `json:"net_Salary"`
}
