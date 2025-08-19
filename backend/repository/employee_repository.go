package repository

import (
	"backend/configs"
	models "backend/model"
)

func CreateEmployee(Employee *models.Employee) error {
	return configs.DB.Create(Employee).Error
}
func GetAllEmployee() ([]models.Employee, error) {
	var employees []models.Employee
	err := configs.DB.Find(&employees).Error
	return employees, err
}

func UpdateEmployee(Employee *models.Employee) error {
	return configs.DB.Save(Employee).Error
}
func GetEmployeeByID(id uint) (*models.Employee, error) {
	var employee models.Employee
	err := configs.DB.First(&employee, id).Error
	if err != nil {
		return nil, err
	}
	return &employee, nil
}
