package main

import (
	"github.com/ProgrammerPeasant/order-control/config"
	"github.com/ProgrammerPeasant/order-control/routes"
	"log"
)

func main() {
	// ПОдгрузка переменных окружения
	config.LoadEnv()
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
	// Инициализируем руты gin
	r := routes.InitRoutes(db)

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
