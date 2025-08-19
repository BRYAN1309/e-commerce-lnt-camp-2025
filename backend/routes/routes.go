package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	api := r.Group("/admin/products")
	{
		api.POST("", handlers.CreateProduct)
		api.GET("", handlers.GetProducts)
		api.PUT("/:id", handlers.UpdateProduct)
	}
}
func AuthRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}
}
func EmployeeRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	{
		api.POST("/users", handlers.CreateEmployee)
		api.GET("/users", handlers.GetAllEmployee)
		api.PUT("/users/:id", handlers.UpdateEmployee)
	}
}
func DashboardRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	{
		api.GET("/dashboard", handlers.GetDashboardStats)
	}
}
func LandingPage(r *gin.Engine) {
	api := r.Group("/products")
	{
		api.GET("/latest", handlers.GetNewestProduct)
		api.GET("/available", handlers.GetAvailableProducts)
	}
}
func Excel(r *gin.Engine) {
	api := r.Group("/admin/products")
	{
		api.GET("export", handlers.ExportProductsToExcel)
	}
}
