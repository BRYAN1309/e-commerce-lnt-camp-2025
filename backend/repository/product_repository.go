package repository

import (
	"backend/configs"
	models "backend/model"
)

func CreateProduct(product *models.Product) error {
	return configs.DB.Create(product).Error
}

func GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	err := configs.DB.Find(&products).Error
	return products, err
}

func GetProductByID(id uint) (*models.Product, error) {
	var product models.Product
	err := configs.DB.First(&product, id).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func UpdateProduct(product *models.Product) error {
	return configs.DB.Save(product).Error
}

func DeleteProduct(product *models.Product) error {
	return configs.DB.Delete(product).Error
}
