package routes

import (
	"github.com/ProgrammerPeasant/order-control/controllers"
	"github.com/ProgrammerPeasant/order-control/middlewares"
	"github.com/ProgrammerPeasant/order-control/repositories"
	"github.com/ProgrammerPeasant/order-control/services"
	"github.com/ProgrammerPeasant/order-control/utils"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

func InitRoutes(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	metrics := utils.NewMetrics()

	r.Use(middlewares.PrometheusMiddleware(metrics))

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// репозитории
	userRepo := repositories.NewUserRepository(db, metrics)
	companyRepo := repositories.NewCompanyRepository(db)
	estimateRepo := repositories.NewEstimateRepository(db)
	joinRequestRepo := repositories.NewJoinRequestRepository(db)

	// сервисы
	companyService := services.NewCompanyService(companyRepo)
	estimateService := services.NewEstimateService(estimateRepo)
	joinRequestService := services.NewJoinRequestService(joinRequestRepo, companyRepo, userRepo, metrics)
	userService := services.NewUserService(userRepo, joinRequestService, metrics)

	// контроллеры
	authController := controllers.NewAuthController(userService, metrics)
	companyController := controllers.NewCompanyController(companyService)
	estimateController := controllers.NewEstimateController(estimateService)
	joinRequestController := controllers.NewJoinRequestController(joinRequestService)

	// Маршруты аутентификации
	auth := r.Group("/api")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/admin/register", middlewares.AuthMiddleware(), middlewares.RoleMiddleware("users:create"), authController.AdminRegister)
	}

	// Маршруты для компаний доступны только авторизованным
	companies := r.Group("/api/v1/companies")
	companies.Use(middlewares.AuthMiddleware())
	{
		// Например, только ADMIN может создавать компании (общее право доступа "companies:create")
		companies.POST("/", middlewares.RoleMiddleware("companies:create"), companyController.CreateCompany) // RoleMiddleware для общих прав

		companies.GET("/:id", companyController.GetCompany) // Доступно всем авторизованным

		// Обновлять и удалять компанию может только ADMIN или MANAGER своей компании
		companies.PUT("/:id", middlewares.CompanyRoleMiddleware(db, "company", "companies:update"), companyController.UpdateCompany)    // CompanyRoleMiddleware
		companies.DELETE("/:id", middlewares.CompanyRoleMiddleware(db, "company", "companies:delete"), companyController.DeleteCompany) // CompanyRoleMiddleware

		// Маршруты для одобрения/отклонения запросов на присоединение (только для менеджеров)
		companies.GET("/join-request", middlewares.CompanyRoleMiddleware(db, "join_request", "join_request:read"), joinRequestController.GetPendingJoinRequests)
		companies.POST("/join-request/approve", middlewares.CompanyRoleMiddleware(db, "join_request", "join_request:accept"), joinRequestController.ApproveJoinRequest)
		companies.POST("/join-request/reject", middlewares.CompanyRoleMiddleware(db, "join_request", "join_request:reject"), joinRequestController.RejectJoinRequest)
	}

	estimateGroup := r.Group("/api/v1/estimates")
	estimateGroup.Use(middlewares.AuthMiddleware()) //  AuthMiddleware для проверки авторизации, CompanyRoleMiddleware для контекстных прав
	{
		// Создавать смету может MANAGER своей компании или ADMIN (контекстно-зависимые права "estimates:create")
		estimateGroup.POST("/", middlewares.CompanyRoleMiddleware(db, "estimate", "estimates:create"), estimateController.CreateEstimate)

		estimateGroup.GET("/:id", estimateController.GetEstimateByID) //  Чтение доступно всем авторизованным
		estimateGroup.PUT("/:id", middlewares.CompanyRoleMiddleware(db, "estimate", "estimates:update"), estimateController.UpdateEstimate)
		estimateGroup.DELETE("/:id", middlewares.CompanyRoleMiddleware(db, "estimate", "estimates:delete"), estimateController.DeleteEstimate)

		estimateGroup.GET("/company", middlewares.RoleMiddleware("companies:create"), estimateController.GetEstimateByCompany)

		// Новый маршрут для получения собственных смет пользователя
		estimateGroup.GET("/my", estimateController.GetMyEstimates)

		estimateGroup.GET("/:id/export/excel", estimateController.ExportEstimateToExcel)
	}

	api := r.Group("/api")
	{
		api.GET("/test", func(c *gin.Context) {
			c.String(http.StatusOK, "Test endpoint is working!")
		})
	}

	return r
}
