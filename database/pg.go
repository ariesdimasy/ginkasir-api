package database

import (
	"ginkasir/config"
	"ginkasir/models"

	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitPgDB() {

	dbUser := config.GetEnv("DB_USER", "postgres")
	dbPass := config.GetEnv("DB_PASS", "999999")
	dbHost := config.GetEnv("DB_HOST", "localhost")
	dbPort := config.GetEnv("DB_PORT", "5432")
	dbName := config.GetEnv("DB_NAME", "ginkasir_db")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s TimeZone=Asia/Jakarta",
		dbHost, dbUser, dbPass, dbName, dbPort,
	)

	var err error

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database : ", err)
	}

	fmt.Println("Database connect sucessfully")

	err = DB.AutoMigrate(
		&models.Category{},
		&models.Product{},
	)

	if err != nil {
		log.Fatal("Failed to migrate database : ", err)
	}

	fmt.Println("Database migrated successfully")

}
