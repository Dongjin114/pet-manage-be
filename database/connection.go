// package database

// import (
// 	"log"
// 	"pet-manage-be/models"

// 	"gorm.io/driver/sqlite"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// // ConnectDB initializes database connection
// func ConnectDB() {
// 	var err error
// 	DB, err = gorm.Open(sqlite.Open("pet_management.db"), &gorm.Config{})
// 	if err != nil {
// 		log.Fatal("Failed to connect to database:", err)
// 	}

// 	log.Println("Database connected successfully")
// }

// // MigrateDB runs database migrations
// func MigrateDB() {
// 	err := DB.AutoMigrate(
// 		&models.Pet{},
// 		&models.Owner{},
// 		&models.MealItem{},
// 		&models.MealHistory{},
// 		&models.MealUnit{},
// 	)
// 	if err != nil {
// 		log.Fatal("Failed to migrate database:", err)
// 	}

// 	log.Println("Database migration completed")
// }
