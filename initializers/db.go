package initializers

import (
	"fmt"
	"IMULIB/models"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB(config *Config) {
	var err error

	// Set the SQLite3 database path
	// dbPath := config.DBPath

	con, _ := LoadConfig(".")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kolkata",
		con.DBHost, con.DBUserName, con.DBUserPassword, con.DBName, con.DBPort,
	)

	// Open the PostgreSQL connection
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the Database! \n", err.Error())
		os.Exit(1)
	}

	log.Println("Running Migrations")
	err = DB.AutoMigrate(&models.User{}, &models.Book{})
	if err != nil {
		log.Fatal("Migration Failed:  \n", err.Error())
		os.Exit(1)
	}

	log.Println("🚀 Connected Successfully to the Database")
}
