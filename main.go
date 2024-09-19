package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Cameron-Johnstone/Noma_Travel/models"
)

var db *gorm.DB

func initDB(config *Config) {
	var err error
	db, err = gorm.Open(postgres.Open(config.GetDBConnString()), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto Migrate the models
	err = db.AutoMigrate(&models.User{}, &models.TravelPlan{}, &models.Location{}, &models.PlanLocation{}, &models.Message{})
	if err != nil {
		log.Fatal("Failed to auto migrate models:", err)
	}
}

func main() {
	// Load configuration
	config := LoadConfig()

	// Initialize database connection
	initDB(config)

	// Initialize router
	r := gin.Default()

	// Define routes
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to Noma!",
		})
	})

	// Start server
	r.Run(":8080")
}
