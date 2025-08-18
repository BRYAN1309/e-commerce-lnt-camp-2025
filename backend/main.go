package main

import (
	"backend/configs"
	"backend/databases"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Setup database connection
	configs.SetUpDatabase()

	// 2. Run AutoMigrate
	databases.AutoMigrate()

	// 3. Setup Gin
	r := gin.Default()

	// 4. Register Routes
	routes.ProductRoutes(r)

	// 5. Start server
	log.Println("🚀 Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
