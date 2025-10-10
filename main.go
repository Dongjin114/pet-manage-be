package main

import (
	"log"
	"net/http"

	"pet-manage-be/handlers"
	"pet-manage-be/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Println("No .env file found")
	// }

	status1 := models.StatusActive
	status2 := models.StatusInactive
	log.Println("status log", status1, status2)

	// // Initialize database
	// database.ConnectDB()
	// database.MigrateDB()

	// // Create Gin router
	// r := gin.Default()

	// // Middleware
	// r.Use(gin.Logger())
	// r.Use(gin.Recovery())
	// r.Use(corsMiddleware())

	// // Routes
	// setupRoutes(r)

	// // Start server
	// port := ":8080"
	// log.Printf("Server starting on port %s", port)
	// if err := r.Run(port); err != nil {
	// 	log.Fatal("Failed to start server:", err)
	// }
}

// CORS middleware
func corsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// Setup all routes
func setupRoutes(r *gin.Engine) {
	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "ok",
			"message": "Pet Management API is running",
		})
	})

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// Pet routes
		pets := v1.Group("/pets")
		{
			pets.GET("", handlers.GetPets)
			pets.GET("/:id", handlers.GetPet)
			pets.POST("", handlers.CreatePet)
			pets.PUT("/:id", handlers.UpdatePet)
			pets.DELETE("/:id", handlers.DeletePet)
		}

		// Owner routes
		owners := v1.Group("/owners")
		{
			owners.GET("", handlers.GetOwners)
			owners.GET("/:id", handlers.GetOwner)
			owners.POST("", handlers.CreateOwner)
			owners.PUT("/:id", handlers.UpdateOwner)
			owners.DELETE("/:id", handlers.DeleteOwner)
		}

		// Meal routes
		meals := v1.Group("/meals")
		{
			meals.GET("/types", handlers.GetMealTypes)
			meals.GET("/items", handlers.GetMealItems)
			meals.GET("/items/:id", handlers.GetMealItem)
			meals.POST("/items", handlers.CreateMealItem)
			meals.PUT("/items/:id", handlers.UpdateMealItem)
			meals.DELETE("/items/:id", handlers.DeleteMealItem)
		}
	}
}
