package main

import (
	"backend/configs"
	"backend/databases"
	"backend/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found")
	}
	// 1. Setup database connection
	configs.SetUpDatabase()

	// 2. Run AutoMigrate
	databases.AutoMigrate()

	// 3. Setup Gin
	r := gin.Default()

	// 4. Register Routes
	routes.ProductRoutes(r)
	routes.AuthRoutes(r)
	// 5. Start server
	log.Println("🚀 Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
