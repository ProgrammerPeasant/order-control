package middlewares

import (
	"net/http"
	"order-control/models"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(allowedRoles ...models.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		roleIf, exists := ctx.Get("role")
		if !exists {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Роль не найдена"})
			ctx.Abort()
			return
		}

		userRole := roleIf.(models.Role)
		// Проверяем, есть ли userRole в списке allowedRoles
		authorized := false
		for _, r := range allowedRoles {
			if r == userRole {
				authorized = true
				break
			}
		}

		if !authorized {
			ctx.JSON(http.StatusForbidden, gin.H{"error": "Недостаточно прав"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
