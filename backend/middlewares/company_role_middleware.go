package middlewares

import (
	"net/http"
	"strconv"

	"github.com/ProgrammerPeasant/order-control/config"
	"github.com/ProgrammerPeasant/order-control/repositories" // Импортируй репозитории
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// CompanyRoleMiddleware проверяет права доступа с учетом принадлежности к компании.
// resourceType - тип ресурса ("estimate" или "company")
// permissionRequired - необходимое право доступа (например, "estimates:update")
func CompanyRoleMiddleware(db *gorm.DB, resourceType string, permissionRequired string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleNameIf, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Роль пользователя не найдена"})
			ctx.Abort()
			return
		}

		userRoleName, ok := roleNameIf.(string)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат роли пользователя"})
			ctx.Abort()
			return
		}

		companyIDIf, exists := ctx.Get("companyID") // Получаем companyID пользователя из контекста
		if !exists {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Company ID пользователя не найден в контексте"})
			ctx.Abort()
			return
		}

		userCompanyID, ok := companyIDIf.(uint)
		if !ok {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат Company ID пользователя"})
			ctx.Abort()
			return
		}

		rolesConfig := config.GetRolesConfig()
		role, roleExists := rolesConfig.Roles[userRoleName]
		if !roleExists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Роль не определена в конфигурации"})
			ctx.Abort()
			return
		}

		// Проверяем общее право доступа (например, "estimates:update") - первый уровень защиты
		hasGeneralPermission := false
		for _, rolePermission := range role.Permissions {
			if rolePermission == permissionRequired {
				hasGeneralPermission = true
				break
			}
		}

		if !hasGeneralPermission {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно общих прав доступа"})
			ctx.Abort()
			return
		}

		// Дальнейшая проверка зависит от типа ресурса (estimate или company) и роли (MANAGER)
		if userRoleName == "MANAGER" {
			if resourceType == "join-request" {
				switch permissionRequired {
				case "join_requests:read":
					// Менеджер может читать запросы своей компании
					ctx.Next()
					return
				case "join_requests:accept", "join_requests:reject":
					ctx.Next()
					return
				default:
					ctx.JSON(http.StatusForbidden, gin.H{"error": "Неизвестное право доступа для join-request"})
					ctx.Abort()
					return
				}
			} else if resourceType == "estimate" || resourceType == "company" { // Исправленная структура if-else if
				if permissionRequired == "estimates:create" || permissionRequired == "companies:create" {
					// Для создания сметы или компании проверяем соответствие CompanyID из контекста
					_, companyIDExists := ctx.Get("companyID")
					if !companyIDExists {
						ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Company ID пользователя не найден в контексте"})
						ctx.Abort()
						return
					}
					// На данном этапе мы предполагаем, что контроллер/сервис будет устанавливать
					// CompanyID создаваемого ресурса равным userCompanyID из контекста.
					// Middleware здесь проверяет лишь право пользователя на создание
					// в контексте своей компании (наличие companyID в контексте).
					ctx.Next()
					return
				} else {
					resourceIDStr := ctx.Param("id")
					if resourceIDStr == "" {
						ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID ресурса не указан"})
						ctx.Abort()
						return
					}

					resourceID, err := strconv.ParseUint(resourceIDStr, 10, 64)
					if err != nil {
						ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат ID ресурса"})
						ctx.Abort()
						return
					}

					var resourceCompanyID uint

					switch resourceType {
					case "estimate":
						estimateRepo := repositories.NewEstimateRepository(db)
						estimate, err := estimateRepo.GetByID(int64(resourceID))
						if err != nil || estimate == nil {
							ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении сметы или смета не найдена"})
							ctx.Abort()
							return
						}
						resourceCompanyID = estimate.CompanyID
					case "company":
						companyRepo := repositories.NewCompanyRepository(db)
						company, err := companyRepo.GetCompanyByID(uint(resourceID))
						if err != nil || company == nil {
							ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении компании или компания не найдена"})
							ctx.Abort()
							return
						}
						resourceCompanyID = company.ID
					}

					if resourceCompanyID != userCompanyID {
						ctx.JSON(http.StatusForbidden, gin.H{"error": "Нет прав доступа к ресурсам другой компании"})
						ctx.Abort()
						return
					}
				}
			}
		}

		ctx.Next() // Доступ разрешен, если все проверки пройдены
	}
}
