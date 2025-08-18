package handlers

import (
	model "backend/model"
	"backend/repository"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateProduct(c *gin.Context) {
	name := c.PostForm("name")
	description := c.PostForm("description")
	price := c.PostForm("price_cents")
	stock := c.PostForm("stock")
	responsibleUserIDStr := c.PostForm("responsible_user_id")

	// Handle image upload
	file, err := c.FormFile("image")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Image is required"})
		return
	}

	// Save image
	filename := uuid.New().String() + filepath.Ext(file.Filename)
	savePath := "uploads/" + filename
	if err := c.SaveUploadedFile(file, savePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload image"})
		return
	}

	// Convert responsible_user_id to *uint
	var responsibleUserID *uint
	if responsibleUserIDStr != "" {
		if v, err := strconv.ParseUint(responsibleUserIDStr, 10, 64); err == nil {
			u := uint(v)
			responsibleUserID = &u
		}
	}

	// Convert price to int64
	var priceCents int64
	if v, err := strconv.ParseInt(price, 10, 64); err == nil {
		priceCents = v
	}

	// Convert stock to int
	var stockInt int
	if v, err := strconv.Atoi(stock); err == nil {
		stockInt = v
	}

	// Save image URL as *string
	imageURL := savePath

	// Build product
	product := model.Product{
		Name:              name,
		Description:       description,
		PriceCents:        priceCents,
		Stock:             stockInt,
		ImageURL:          &imageURL,
		ResponsibleUserID: responsibleUserID,
	}

	// Save to DB
	if err := repository.CreateProduct(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
		return
	}

	c.JSON(http.StatusOK, product)
}

// GET /api/products
func GetProducts(c *gin.Context) {
	products, err := repository.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

// GET /api/products/:id
func GetProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repository.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// PUT /api/products/:id
func UpdateProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repository.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}

	if name := c.PostForm("name"); name != "" {
		product.Name = name
	}
	if desc := c.PostForm("description"); desc != "" {
		product.Description = desc
	}
	if s := c.PostForm("price_cents"); s != "" {
		if v, err := strconv.ParseInt(s, 10, 64); err == nil {
			product.PriceCents = v
		}
	}
	if s := c.PostForm("stock"); s != "" {
		if v, err := strconv.Atoi(s); err == nil {
			product.Stock = v
		}
	}
	if s := c.PostForm("responsible_user_id"); s != "" {
		if v, err := strconv.ParseUint(s, 10, 64); err == nil {
			u := uint(v)
			product.ResponsibleUserID = &u
		}
	}

	// handle image update
	if file, err := c.FormFile("image"); err == nil {
		ext := strings.ToLower(filepath.Ext(file.Filename))
		filename := fmt.Sprintf("%s_%d%s", uuid.NewString(), time.Now().UnixNano(), ext)
		path := filepath.Join("backend/uploads", filename)

		if saveErr := c.SaveUploadedFile(file, path); saveErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": saveErr.Error()})
			return
		}
		// delete old image
		if product.ImageURL != nil {
			_ = os.Remove(*product.ImageURL)
		}
		product.ImageURL = &path
	}

	if err := repository.UpdateProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update product"})
		return
	}
	c.JSON(http.StatusOK, product)
}

// DELETE /api/products/:id
func DeleteProduct(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	product, err := repository.GetProductByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "product not found"})
		return
	}
	// remove image file if exists
	if product.ImageURL != nil {
		_ = os.Remove(*product.ImageURL)
	}
	if err := repository.DeleteProduct(product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete product"})
		return
	}
	c.Status(http.StatusNoContent)
}
