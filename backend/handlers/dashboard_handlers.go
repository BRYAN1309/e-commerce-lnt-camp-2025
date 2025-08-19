package handlers

import (
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetDashboardStats(c *gin.Context) {
	employeeCount, err := repository.CountEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count employees"})
		return
	}

	activeEmployeeCount, err := repository.CountActiveEmployees()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count active employees"})
		return
	}

	productCount, err := repository.CountProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count products"})
		return
	}

	availableProductCount, err := repository.CountAvailableProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to count available products"})
		return
	}

	latestProducts, err := repository.GetLatestProducts(5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch latest products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"employee_total":        employeeCount,
		"employee_active_total": activeEmployeeCount,
		"product_total":         productCount,
		"product_available":     availableProductCount,
		"latest_products":       latestProducts,
	})
}
