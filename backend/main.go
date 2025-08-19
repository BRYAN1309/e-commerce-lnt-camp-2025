package main

import (
	"backend/configs"
	"backend/databases"
	"backend/databases/seeders"
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

	// 2. Run AutoMigrate FIRST - Create tables before seeding
	databases.AutoMigrate()

	// 3. THEN run seeders - Insert data after tables are created
	seeders.SeederProducts()
	seeders.SeederUser()
	seeders.SeedersEmployee()

	// 4. Setup Gin
	r := gin.Default()

	// 5. Register Routes
	routes.ProductRoutes(r)
	routes.AuthRoutes(r)
	routes.EmployeeRoutes(r)
	routes.DashboardRoutes(r)
	routes.LandingPage(r)
	routes.Excel(r)
	// 6. Start server
	log.Println("🚀 Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal("❌ Failed to start server: ", err)
	}
}
