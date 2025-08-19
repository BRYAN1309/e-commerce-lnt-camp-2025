package seeders

import (
	"backend/configs"
	models "backend/model"
	"errors"

	"gorm.io/gorm"
)

func SeederProducts() {
	err := configs.DB.First(&models.Product{}).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		var products []models.Product

		image1 := "uploads/product1.jpg"
		image2 := "uploads/product2.jpg"
		image3 := "uploads/product3.jpg"

		userID := uint(1)

		products = append(products,
			models.Product{
				Name:              "Laptop Gaming",
				Description:       "High-performance gaming laptop with RTX graphics",
				PriceCents:        12999900,
				Stock:             15,
				ImageURL:          &image1,
				ResponsibleUserID: &userID,
			},
			models.Product{
				Name:              "Wireless Headphones",
				Description:       "Noise-cancelling Bluetooth headphones",
				PriceCents:        29900,
				Stock:             42,
				ImageURL:          &image2,
				ResponsibleUserID: &userID,
			},
			models.Product{
				Name:              "Smartphone Pro",
				Description:       "Latest smartphone with advanced camera system",
				PriceCents:        89900,
				Stock:             28,
				ImageURL:          &image3,
				ResponsibleUserID: &userID,
			},
			models.Product{
				Name:              "Mechanical Keyboard",
				Description:       "RGB mechanical keyboard with cherry MX switches",
				PriceCents:        15900,
				Stock:             35,
				ImageURL:          nil,
				ResponsibleUserID: &userID,
			},
			models.Product{
				Name:              "Gaming Mouse",
				Description:       "Ergonomic gaming mouse with customizable DPI",
				PriceCents:        7900,
				Stock:             50,
				ImageURL:          nil,
				ResponsibleUserID: &userID,
			})

		if err := configs.DB.Create(&products).Error; err != nil {
			println("Error seeding products:", err.Error())
		} else {
			println("Successfully seeded products")
		}
	}
}
