package service

import (
	"sukvij/employment/model"
	"sukvij/employment/repository"

	"gorm.io/gorm"
)

type Service struct {
	Db       *gorm.DB
	Employee *model.Employee
}

func (service *Service) CreateEmployee() (*[]model.Employee, error) {
	repo := &repository.Repository{Db: service.Db, Employee: service.Employee}
	return repo.CreateEmployee()
}

func (service *Service) GetEmployee() (*[]model.Employee, error) {
	repo := &repository.Repository{Db: service.Db, Employee: service.Employee}
	return repo.GetEmployee()
}

func (service *Service) GetEmployeeById(id uint) (*model.Employee, error) {
	repo := &repository.Repository{Db: service.Db, Employee: service.Employee}
	return repo.GetEmployeeById(id)
}
func (service *Service) UpdateEmployee() {

}

func (service *Service) DeleteEmployee(id uint) error {
	repo := &repository.Repository{Db: service.Db, Employee: service.Employee}
	return repo.DeleteEmployee(id)
}
