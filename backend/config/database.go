package config

import (
	"fmt"
	"github.com/ProgrammerPeasant/order-control/models"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Warning: No .env file found")
	}
}

// InitDB инициализирует подключение к mysql через gorm
func InitDB() (*gorm.DB, error) {

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" || dbPort == "" {
		log.Println("Некоторые переменные окружения для базы данных не заданы")
		return nil, fmt.Errorf("missing env vars for DB connection")
	}
	log.Printf("Подключение...")
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbName, dbPass)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	// Можно включать/выключать логирование SQL-запросов
	db.LogMode(true)

	db.AutoMigrate(
		&models.User{},
		&models.Company{},
		&models.Estimate{},
		&models.EstimateItem{})

	return db, nil
}
