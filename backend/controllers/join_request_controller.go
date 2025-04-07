package controllers

import (
	"net/http"

	"github.com/ProgrammerPeasant/order-control/services"
	"github.com/gin-gonic/gin"
)

type JoinRequestController struct {
	joinRequestService services.JoinRequestService
}

func NewJoinRequestController(joinRequestService services.JoinRequestService) *JoinRequestController {
	return &JoinRequestController{joinRequestService: joinRequestService}
}

// GetPendingJoinRequests
// @Summary Получить запросы на присоединение к компании (только для менеджеров)
// @Description Возвращает список запросов на присоединение к компании текущего менеджера. Требуется роль менеджера.
// @Tags Companies
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.JoinRequest "Список запросов на присоединение"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Доступ запрещен. Требуется роль менеджера."
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies/join-request [get]
func (c *JoinRequestController) GetPendingJoinRequests(ctx *gin.Context) {
	companyIDInterface, exists := ctx.Get("companyID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}
	companyID, ok := companyIDInterface.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	joinRequests, err := c.joinRequestService.GetPendingJoinRequests(companyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, joinRequests)
}

// ApproveJoinRequest
// @Summary Одобрить запрос на присоединение к компании (только для менеджеров)
// @Description Одобряет запрос на присоединение пользователя к компании. Требуется роль менеджера.
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body controllers.ApproveRejectRequest true "ID пользователя для одобрения"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H{message=string} "Запрос одобрен"
// @Failure 400 {object} gin.H "Неверный запрос"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Доступ запрещен. Требуется роль менеджера."
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies/join-request/approve [post]
func (c *JoinRequestController) ApproveJoinRequest(ctx *gin.Context) {
	companyIDInterface, exists := ctx.Get("companyID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}
	companyID, ok := companyIDInterface.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	var request ApproveRejectRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	err := c.joinRequestService.ApproveJoinRequest(request.UserID, companyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Запрос одобрен"})
}

// RejectJoinRequest
// @Summary Отклонить запрос на присоединение к компании (только для менеджеров)
// @Description Отклоняет запрос на присоединение пользователя к компании. Требуется роль менеджера.
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body controllers.ApproveRejectRequest true "ID пользователя для отклонения"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H{message=string} "Запрос отклонен"
// @Failure 400 {object} gin.H "Неверный запрос"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Доступ запрещен. Требуется роль менеджера."
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies/join-request/reject [post]
func (c *JoinRequestController) RejectJoinRequest(ctx *gin.Context) {
	companyIDInterface, exists := ctx.Get("companyID")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Не авторизован"})
		return
	}
	companyID, ok := companyIDInterface.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
		return
	}

	var request ApproveRejectRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный запрос"})
		return
	}

	err := c.joinRequestService.RejectJoinRequest(request.UserID, companyID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Запрос отклонен"})
}
