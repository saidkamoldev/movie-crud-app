// package repository

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// 	"log"
// )

// var DB *gorm.DB

// // ConnectDB bazaga ulanishni yaratadi
//
//	func ConnectDB() *gorm.DB {
//		dsn := "user=youruser password=yourpassword dbname=yourdb port=5432 sslmode=disable"
//		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//		if err != nil {
//			log.Fatalf("failed to connect to database: %v", err)
//		}
//		DB = db
//		return DB
//	}
package repository

import (
	"log"
	"movie-crud-app/internal/repository/models"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB
var DB *gorm.DB

func ConnectDB() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	dsn := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = db.AutoMigrate(&models.Movie{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	DB = db
	return DB
}
