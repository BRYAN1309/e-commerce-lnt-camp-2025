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
		api.GET("/:id", handlers.GetProduct)
		api.PUT("/:id", handlers.UpdateProduct)
		api.DELETE("/:id", handlers.DeleteProduct)
	}
}
func AuthRoutes(r *gin.Engine) {
	api := r.Group("/admin")
	{
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
	}
}
