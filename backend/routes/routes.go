package routes

import (
	"backend/handlers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(r *gin.Engine) {
	api := r.Group("/api/products")
	{
		api.POST("", handlers.CreateProduct)
		api.GET("", handlers.GetProducts)
		api.GET("/:id", handlers.GetProduct)
		api.PUT("/:id", handlers.UpdateProduct)
		api.DELETE("/:id", handlers.DeleteProduct)
	}
}
