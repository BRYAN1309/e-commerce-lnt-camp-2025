package handlers

import (
	models "backend/model"
	"backend/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateEmployee(c *gin.Context) {
	var input struct {
		Name        string `json:"name" binding:"required"`
		Email       string `json:"email" binding:"required"`
		Telelephone int64  `json:"telephone" binding:"required"`
		Status      string `json:"status" binding:"required,oneof=aktif tidak aktif"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	Employee := models.Employee{
		Name:        input.Name,
		Email:       input.Email,
		Telelephone: input.Telelephone,
		Status:      input.Status,
	}
	if err := repository.CreateEmployee(&Employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create employee"})
		return
	}
	c.JSON(http.StatusOK, Employee)
}

func GetAllEmployee(c *gin.Context) {
	Employee, err := repository.GetAllEmployee()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch employee"})
		return
	}
	c.JSON(http.StatusOK, Employee)
}

func UpdateEmployee(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// First, get the existing employee
	employee, err := repository.GetEmployeeByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	// Update fields if provided
	if name := c.PostForm("name"); name != "" {
		employee.Name = name
	}
	if email := c.PostForm("email"); email != "" {
		employee.Email = email
	}
	if telephone := c.PostForm("telephone"); telephone != "" {
		// Convert string to int64
		telInt, err := strconv.ParseInt(telephone, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid telephone number"})
			return
		}
		employee.Telelephone = telInt
	}
	if status := c.PostForm("status"); status != "" {
		// Validate status value
		if status != "aktif" && status != "tidak aktif" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Status must be 'aktif' or 'tidak aktif'"})
			return
		}
		employee.Status = status
	}

	// Save the updated employee
	if err := repository.UpdateEmployee(employee); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update employee"})
		return
	}

	c.JSON(http.StatusOK, employee)
}
