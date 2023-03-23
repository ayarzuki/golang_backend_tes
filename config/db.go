package config

import (
	"fmt"
	"os"

	"goproduct/models"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	_ "github.com/jinzhu/gorm/dialects/mysql" // Untuk MySQL
)

func Connect() *gorm.DB {
	err := godotenv.Load("./handlers/.env")
	if err != nil {
		panic("Failed to load .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbCharset := os.Getenv("DB_CHARSET")
	dbParseTime := os.Getenv("DB_PARSE_TIME")
	dbLoc := os.Getenv("DB_LOC")

	dbURL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, dbCharset, dbParseTime, dbLoc)

	db, err := gorm.Open("mysql", dbURL)

	if err != nil {
		panic("Failed to connect to database!")
	}

	db.AutoMigrate(&models.Product{})

	return db
}
