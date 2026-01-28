package repository

import (
	"sukvij/employment/model"

	"gorm.io/gorm"
)

type Repository struct {
	Db       *gorm.DB
	Employee *model.Employee
}

func (repo *Repository) CreateEmployee() (*[]model.Employee, error) {
	var employee []model.Employee
	err := repo.Db.Create(repo.Employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (repo *Repository) GetEmployee() (*[]model.Employee, error) {
	var employee []model.Employee
	err := repo.Db.Find(&employee).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}

func (repo *Repository) GetEmployeeById() {

}
func (repo *Repository) UpdateEmployee() {

}

func (repo *Repository) DeleteEmployee() {

}
