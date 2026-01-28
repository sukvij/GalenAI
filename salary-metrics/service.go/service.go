package servicego

import (
	"sukvij/salary-metrics/model"

	"gorm.io/gorm"
)

type Service struct {
	Db *gorm.DB
}

func (service *Service) SalaryMetricsCountryWise(country string) (*model.SalaryMetrics, error) {
	var metrics model.SalaryMetrics

	err := service.Db.Raw(`
	SELECT
	    COALESCE(MIN(salary), 0)     AS min_salary,
	    COALESCE(MAX(salary), 0)     AS max_salary,
	    COALESCE(AVG(salary), 0)     AS average_salary
	FROM employees
	WHERE country = ?
`, country).Scan(&metrics).Error
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}

func (service *Service) SalaryMetricsJobTitleWise(title string) (*model.SalaryMetrics, error) {
	var metrics model.SalaryMetrics

	err := service.Db.Raw(`
	SELECT
	    COALESCE(MIN(salary), 0)     AS min_salary,
	    COALESCE(MAX(salary), 0)     AS max_salary,
	    COALESCE(AVG(salary), 0)     AS average_salary
	FROM employees
	WHERE job_title = ?
`, title).Scan(&metrics).Error
	if err != nil {
		return nil, err
	}
	return &metrics, nil
}
