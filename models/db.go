package models

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Get database credentials from environment variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Create a database connection string
	dbConnectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)

	// Open a GORM database connection
	gormDB, err := gorm.Open(mysql.Open(dbConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database:", err.Error())
	}

	// AutoMigrate the necessary tables
	err = gormDB.AutoMigrate(&User_Input{})
	if err != nil {
		log.Fatal("Error automigrating database tables:", err.Error())
	}
	// Assign the GORM DB instance to the global variable
	DB = gormDB

	// Perform any additional setup or configurations here

	fmt.Println("Connected to the database successfully!")
}
