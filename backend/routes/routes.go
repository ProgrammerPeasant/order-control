package routes

import (
	"github.com/ProgrammerPeasant/order-control/controllers"
	"github.com/ProgrammerPeasant/order-control/middlewares"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/ProgrammerPeasant/order-control/services"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// Инициализируем репозитории
	userRepo := repositories.NewUserRepository(db)
	companyRepo := repositories.NewCompanyRepository(db)

	// Инициализируем сервисы
	userService := services.NewUserService(userRepo)
	companyService := services.NewCompanyService(companyRepo)

	// Инициализируем контроллеры
	authController := controllers.NewAuthController(userService)
	companyController := controllers.NewCompanyController(companyService)

	// Маршруты аутентификации
	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
	}

	// Маршруты для компаний (доступны только авторизованным)
	companies := r.Group("/companies")
	companies.Use(middlewares.AuthMiddleware())
	{
		// Например, только ADMIN и MANAGER могут создавать компании
		companies.POST("/", middlewares.RoleMiddleware("ADMIN", "MANAGER"), companyController.CreateCompany)

		// Доступно всем авторизованным
		companies.GET("/:id", companyController.GetCompany)

		// Допустим, обновлять и удалять может только ADMIN
		companies.PUT("/:id", middlewares.RoleMiddleware("ADMIN"), companyController.UpdateCompany)
		companies.DELETE("/:id", middlewares.RoleMiddleware("ADMIN"), companyController.DeleteCompany)
	}

	return r
}
