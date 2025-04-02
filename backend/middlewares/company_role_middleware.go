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
		if userRoleName == "MANAGER" { // Контекстно-зависимая проверка только для менеджеров
			resourceIDStr := ctx.Param("id") // Предполагаем, что ID ресурса передается как параметр пути "id"
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

			// В зависимости от resourceType, получаем CompanyID ресурса из БД
			switch resourceType {
			case "estimate":
				estimateRepo := repositories.NewEstimateRepository(db) // Используем репозиторий estimates
				estimate, err := estimateRepo.GetByID(int64(resourceID))
				if err != nil || estimate == nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении сметы или смета не найдена"})
					ctx.Abort()
					return
				}
				resourceCompanyID = estimate.CompanyID // Получаем CompanyID из сметы
			case "company":
				companyRepo := repositories.NewCompanyRepository(db) // Используем репозиторий company
				company, err := companyRepo.GetCompanyByID(uint(resourceID))
				if err != nil || company == nil {
					ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении компании или компания не найдена"})
					ctx.Abort()
					return
				}
				resourceCompanyID = company.ID // CompanyID компании - это ее ID
			default:
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Неизвестный тип ресурса для контекстной проверки прав"})
				ctx.Abort()
				return
			}

			// Проверяем, принадлежит ли ресурс компании менеджера
			if resourceCompanyID != userCompanyID {
				ctx.JSON(http.StatusForbidden, gin.H{"error": "Нет прав доступа к ресурсам другой компании"}) // Ошибка контекста - ресурс не из компании менеджера
				ctx.Abort()
				return
			}
		}

		ctx.Next() // Доступ разрешен, если все проверки пройдены
	}
}
