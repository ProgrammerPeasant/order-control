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

	estimateRepo := repositories.NewEstimateRepositories(db)
	estimateService := services.NewEstimateService(estimateRepo)
	estimateController := controllers.NewEstimateController(estimateService)

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

	authMiddleware := middlewares.AuthMiddleware()
	roleMiddleware := middlewares.RoleMiddleware("MANAGER", "ADMIN")

	estimateGroup := r.Group("/api/v1/estimates")
	estimateGroup.Use(authMiddleware, roleMiddleware)
	{
		estimateGroup.POST("/", estimateController.CreateEstimate)

		estimateGroup.GET("/:id", estimateController.GetEstimateByID)   // GET /api/v1/estimates/:id - Get estimate by ID
		estimateGroup.PUT("/:id", estimateController.UpdateEstimate)    // PUT /api/v1/estimates/:id - Update estimate by ID
		estimateGroup.DELETE("/:id", estimateController.DeleteEstimate) // DELETE /api/v1/estimates/:id - Delete estimate by ID

		estimateGroup.GET("/company", estimateController.GetEstimateByCompany) // GET /api/v1/estimates/company?company_id=... - Get estimates by Company ID
	}

	return r
}
