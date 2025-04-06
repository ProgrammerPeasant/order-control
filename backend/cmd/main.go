package main

import (
	"fmt"
	swaggerFiles "github.com/swaggo/files"
	"log"

	"github.com/ProgrammerPeasant/order-control/config" // Предполагается, что этот путь правильный
	"github.com/ProgrammerPeasant/order-control/routes" // Предполагается, что этот путь правильный

	"github.com/gin-contrib/cors" // Импортируем middleware CORS для Gin
	ginSwagger "github.com/swaggo/gin-swagger"
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

	// **Добавляем middleware CORS здесь:**
	configCors := cors.DefaultConfig()
	configCors.AllowAllOrigins = true
	configCors.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	configCors.AllowHeaders = []string{"Accept", "Accept-Language", "Content-Type", "Authorization"}

	r.Use(cors.New(configCors))

	_ = ginSwagger.URL("/swagger/index.html") // URL, по которому будет доступен Swagger UI
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Запуск сервера
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
