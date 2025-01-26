package cmd

import (
	"log"
	"myapp/config"
	"myapp/routes"
)

func main() {
	// Инициализируем базу данных
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	// Закрываем базу данных при завершении
	defer func() {
		if err := db.Close(); err != nil {
			log.Printf("Ошибка при закрытии БД: %v", err)
		}
	}()

	// Инициализируем роуты (Gin)
	r := routes.InitRoutes(db)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
