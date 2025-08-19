package routes

import (
	"backend/handlers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	api := r.Group("/admin/products")
	api.Use(middleware.AuthMiddleware()) // protect products with JWT
	{
		api.POST("", handlers.CreateProduct)
		api.GET("", handlers.GetProducts)
		api.PUT("/:id", handlers.UpdateProduct)
	}
}

func AuthRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	{
		// Auth does not need middleware
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}
}

func EmployeeRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	api.Use(middleware.AuthMiddleware()) // protect employees with JWT
	{
		api.POST("/users", handlers.CreateEmployee)
		api.GET("/users", handlers.GetAllEmployee)
		api.PUT("/users/:id", handlers.UpdateEmployee)
	}
}

func DashboardRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	api.Use(middleware.AuthMiddleware()) // protect dashboard with JWT
	{
		api.GET("/dashboard", handlers.GetDashboardStats)
	}
}

func LandingPage(r *gin.Engine) {
	api := r.Group("/products")
	{
		// Landing page is public (no JWT needed)
		api.GET("/latest", handlers.GetNewestProduct)
		api.GET("/available", handlers.GetAvailableProducts)
	}
}

func Excel(r *gin.Engine) {
	api := r.Group("/admin/products")
	api.Use(middleware.AuthMiddleware()) // protect export with JWT
	{
		api.GET("/export", handlers.ExportProductsToExcel)
	}
}
