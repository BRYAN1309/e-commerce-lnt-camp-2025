package repository

import (
	"backend/configs"
	models "backend/model"
)

func CountEmployees() (int64, error) {
	var count int64
	err := configs.DB.Model(&models.Employee{}).Count(&count).Error
	return count, err
}

func CountActiveEmployees() (int64, error) {
	var count int64
	err := configs.DB.Model(&models.Employee{}).Where("status = ?", "aktif").Count(&count).Error
	return count, err
}

func CountProducts() (int64, error) {
	var count int64
	err := configs.DB.Model(&models.Product{}).Count(&count).Error
	return count, err
}

func CountAvailableProducts() (int64, error) {
	var count int64
	err := configs.DB.Model(&models.Product{}).Where("stock > ?", 0).Count(&count).Error
	return count, err
}

func GetLatestProducts(limit int) ([]models.Product, error) {
	var products []models.Product
	err := configs.DB.Order("created_at DESC").Limit(limit).Find(&products).Error
	return products, err
}

func GetAvailableProducts() ([]models.Product, error) {
	var products []models.Product
	err := configs.DB.Where("stock > ?", 0).Order("created_at DESC").Find(&products).Error
	return products, err
}
