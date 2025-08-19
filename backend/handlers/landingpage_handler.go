package handlers

import (
	"backend/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetNewestProduct(c *gin.Context) {
	latestProducts, err := repository.GetLatestProducts(5)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch latest products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"latest_products": latestProducts,
	})
}
func GetAvailableProducts(c *gin.Context) {
	products, err := repository.GetAvailableProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch available products"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"available_products": products,
	})
}
