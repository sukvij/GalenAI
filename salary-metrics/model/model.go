package model

type SalaryMetrics struct {
	MinSalary     float64 `json:"min_salary"`
	MaxSalary     float64 `json:"max_salary"`
	AverageSalary float64 `json:"avg_salary"`
}
