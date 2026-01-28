package service

import (
	"fmt"
	"sukvij/employment/repository"
	"sukvij/salary-calculation/model"

	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func (service *Service) SalaryCalculation(id uint) (*model.Salary, error) {
	var salary model.Salary

	repo := &repository.Repository{Db: service.Db}
	employee, err := repo.GetEmployeeById(id)

	if err != nil {
		return nil, err
	}

	salary.ID = employee.ID
	salary.FullName = employee.FullName
	salary.JobTitle = employee.JobTitle
	salary.Country = employee.Country

	mp := map[string]float64{}
	mp["india"] = 0.10
	mp["united states"] = 0.12

	deductionPercent := float64(0)
	// we just need to calculate salary of employee

	if val, ok := mp[employee.Country]; ok {
		deductionPercent = val
	}

	fmt.Println(deductionPercent, " % h ye", employee.Country)
	salary.GrossSalary = employee.Salary
	salary.DeductionAmount = (employee.Salary * deductionPercent)
	salary.NetSalary = employee.Salary - salary.DeductionAmount
	return &salary, nil
}
