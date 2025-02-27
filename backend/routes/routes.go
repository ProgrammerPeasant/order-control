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
	estimateRepo := repositories.NewEstimateRepositories(db) // Инициализация estimateRepo

	// Инициализируем сервисы
	userService := services.NewUserService(userRepo)
	companyService := services.NewCompanyService(companyRepo)
	estimateService := services.NewEstimateService(estimateRepo) // Инициализация estimateService

	// Инициализируем контроллеры
	authController := controllers.NewAuthController(userService)
	companyController := controllers.NewCompanyController(companyService)
	estimateController := controllers.NewEstimateController(estimateService) // Инициализация estimateController

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
		// Например, только ADMIN может создавать компании (общее право доступа "companies:create")
		companies.POST("/", middlewares.RoleMiddleware("companies:create"), companyController.CreateCompany) // Используем RoleMiddleware для общих прав

		companies.GET("/:id", companyController.GetCompany) // Доступно всем авторизованным

		// Обновлять и удалять компанию может только ADMIN или MANAGER своей компании (контекстно-зависимые права доступа)
		companies.PUT("/:id", middlewares.CompanyRoleMiddleware(db, "company", "companies:update"), companyController.UpdateCompany)    // Используем CompanyRoleMiddleware
		companies.DELETE("/:id", middlewares.CompanyRoleMiddleware(db, "company", "companies:delete"), companyController.DeleteCompany) // Используем CompanyRoleMiddleware
	}

	estimateGroup := r.Group("/api/v1/estimates")
	estimateGroup.Use(middlewares.AuthMiddleware()) //  AuthMiddleware для проверки авторизации, CompanyRoleMiddleware для контекстных прав
	{
		// Создавать смету может MANAGER своей компании или ADMIN (контекстно-зависимые права "estimates:create")
		estimateGroup.POST("/", middlewares.CompanyRoleMiddleware(db, "estimate", "estimates:create"), estimateController.CreateEstimate) // CompanyRoleMiddleware

		estimateGroup.GET("/:id", estimateController.GetEstimateByID)                                                                          //  Чтение доступно всем авторизованным (или можно тоже сделать контекстно-зависимым)
		estimateGroup.PUT("/:id", middlewares.CompanyRoleMiddleware(db, "estimate", "estimates:update"), estimateController.UpdateEstimate)    // CompanyRoleMiddleware
		estimateGroup.DELETE("/:id", middlewares.CompanyRoleMiddleware(db, "estimate", "estimates:delete"), estimateController.DeleteEstimate) // CompanyRoleMiddleware

		estimateGroup.GET("/company", estimateController.GetEstimateByCompany) // Чтение списка смет компании - можно сделать доступным всем авторизованным, или ограничить контекстом
	}

	return r
}
