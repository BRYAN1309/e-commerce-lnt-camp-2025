package seeders

import (
	"backend/configs"
	models "backend/model"
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func SeederUser() {
	err := configs.DB.First(&models.User{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var users []models.User

		// Hash passwords
		password1, _ := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
		password2, _ := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
		password3, _ := bcrypt.GenerateFromPassword([]byte("user123"), bcrypt.DefaultCost)

		users = append(users,
			models.User{
				Name:     "Admin User",
				Email:    "admin@company.com",
				Password: string(password1),
			},
			models.User{
				Name:     "Manager User",
				Email:    "manager@company.com",
				Password: string(password2),
			},
			models.User{
				Name:     "Regular User",
				Email:    "user@company.com",
				Password: string(password3),
			},
			models.User{
				Name:     "John Doe",
				Email:    "john.doe@example.com",
				Password: string(password1),
			},
			models.User{
				Name:     "Jane Smith",
				Email:    "jane.smith@example.com",
				Password: string(password2),
			})

		if err := configs.DB.Create(&users).Error; err != nil {
			println("Error seeding users:", err.Error())
		} else {
			println("Successfully seeded users")
		}
	}
}
