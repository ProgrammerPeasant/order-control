package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// InitDB инициализирует подключение к MySQL через GORM
func InitDB() (*gorm.DB, error) {
	// Пример извлечения параметров для подключения из переменных окружения
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	if dbUser == "" || dbPass == "" || dbHost == "" || dbName == "" || dbPort == "" {
		log.Println("Некоторые переменные окружения для базы данных не заданы")
		return nil, fmt.Errorf("missing env vars for DB connection")
	}

	// Примерная строка подключения
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbPort, dbName)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Можно включать/выключать логирование SQL-запросов
	db.LogMode(true)

	// Миграции (как пример - мигрируем несколько сущностей)
	// db.AutoMigrate(&models.User{}, &models.Company{})

	return db, nil
}
