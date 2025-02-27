package middlewares

import (
	"net/http"

	"github.com/ProgrammerPeasant/order-control/config"
	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedPermissions ...string) gin.HandlerFunc {
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

		rolesConfig := config.GetRolesConfig()

		role, roleExists := rolesConfig.Roles[userRoleName]
		if !roleExists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Роль не определена в конфигурации"})
			ctx.Abort()
			return
		}

		isAllowed := false
		for _, requiredPermission := range allowedPermissions {
			for _, rolePermission := range role.Permissions {
				if rolePermission == requiredPermission {
					isAllowed = true
					break
				}
			}
			if isAllowed {
				break
			}
		}

		if !isAllowed {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав доступа"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
