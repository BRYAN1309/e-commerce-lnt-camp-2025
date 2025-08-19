package seeders

import (
	"backend/configs"
	models "backend/model"
	"errors"

	"gorm.io/gorm"
)

func SeedersEmployee() {
	err := configs.DB.First(&models.Employee{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var employees []models.Employee

		employees = append(employees,
			models.Employee{
				Name:        "John Doe",
				Email:       "john.doe@company.com",
				Telelephone: 628123456789,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "Jane Smith",
				Email:       "jane.smith@company.com",
				Telelephone: 628987654321,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "Robert Johnson",
				Email:       "robert.j@company.com",
				Telelephone: 628555123456,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "Sarah Wilson",
				Email:       "sarah.wilson@company.com",
				Telelephone: 628444789012,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "Michael Brown",
				Email:       "michael.b@company.com",
				Telelephone: 628333456789,
				Status:      "tidak aktif",
			},
			models.Employee{
				Name:        "Emily Davis",
				Email:       "emily.davis@company.com",
				Telelephone: 628222987654,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "David Miller",
				Email:       "david.m@company.com",
				Telelephone: 628111234567,
				Status:      "tidak aktif",
			},
			models.Employee{
				Name:        "Lisa Anderson",
				Email:       "lisa.a@company.com",
				Telelephone: 628999876543,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "James Taylor",
				Email:       "james.t@company.com",
				Telelephone: 628888345678,
				Status:      "aktif",
			},
			models.Employee{
				Name:        "Amanda Clark",
				Email:       "amanda.c@company.com",
				Telelephone: 628777654321,
				Status:      "tidak aktif",
			})

		if err := configs.DB.Create(&employees).Error; err != nil {
			println("Error seeding employees:", err.Error())
		} else {
			println("Successfully seeded employees with status field")
		}
	}
}
