package main

import (
	"fmt"
	_ "github.com/ProgrammerPeasant/order-control/cmd/docs" // docs is generated by Swag CLI, импортируем сгенерированные документы!
	"github.com/ProgrammerPeasant/order-control/config"
	"github.com/ProgrammerPeasant/order-control/routes"
	"github.com/swaggo/files" // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

// @title Order Control API
// @version 1.0
// @description API для управления заказами и сметами в системе Order Control.

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @host localhost:8080
// @BasePath /api

func main() {
	// ПОдгрузка переменных окружения
	config.LoadEnv()

	rolesConfig, err := config.LoadRolesConfig("./config/roles.yaml") // Путь к roles.yaml
	if err != nil {
		log.Fatalf("Failed to load roles configuration: %v", err)
	}
	fmt.Printf("Roles configuration loaded: %+v\n", rolesConfig)

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

	// url := ginSwagger.URL("/swagger/index.html") // URL, по которому будет доступен Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
