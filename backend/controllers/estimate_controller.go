package controllers

import (
	"net/http"
	"strconv"

	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/services"
	"github.com/gin-gonic/gin"
)

type EstimateController struct {
	estimateService *services.EstimateService
}

func NewEstimateController(service *services.EstimateService) *EstimateController {
	return &EstimateController{estimateService: service}
}

// CreateEstimate
// @Summary Создать новую смету
// @Description Создает новую смету на основе данных запроса. Доступно только менеджерам и администраторам своей компании.
// @Tags Estimates
// @Accept json
// @Produce json
// @Param request body models.Estimate true "Данные сметы для создания"
// @Security ApiKeyAuth
// @Success 201 {object} models.Estimate "Смета успешно создана"
// @Failure 400 {object} gin.H "Неверный запрос"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/estimates [post]
func (c *EstimateController) CreateEstimate(ctx *gin.Context) {
	var estimate models.Estimate
	if err := ctx.ShouldBindJSON(&estimate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userIDIf, exists := ctx.Get("userID") // Получаю userID из контекста
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "User ID не найден в контексте"})
		return
	}
	userID, ok := userIDIf.(uint)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Неверный формат User ID в контексте"})
		return
	}
	estimate.CreatedByID = userID

	if err := c.estimateService.CreateEstimate(&estimate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, estimate)
}

// UpdateEstimate
// @Summary Обновить существующую смету
// @Description Обновляет данные существующей сметы по ID. Доступно только менеджерам и администраторам своей компании.
// @Tags Estimates
// @Accept json
// @Produce json
// @Param id path integer true "ID сметы для обновления"
// @Param request body models.Estimate true "Новые данные сметы"
// @Security ApiKeyAuth
// @Success 200 {object} models.Estimate "Смета успешно обновлена"
// @Failure 400 {object} gin.H "Неверный запрос или неверный ID сметы"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Нет прав доступа"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/estimates/{id} [put]
func (c *EstimateController) UpdateEstimate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	estimateID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат Estimate ID"})
		return
	}

	var estimate models.Estimate
	if err := ctx.ShouldBindJSON(&estimate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	estimate.ID = uint(estimateID)

	if err := c.estimateService.UpdateEstimate(&estimate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, estimate)
}

// DeleteEstimate
// @Summary Удалить смету
// @Description Удаляет смету по ID. Доступно только менеджерам и администраторам своей компании.
// @Tags Estimates
// @Param id path integer true "ID сметы для удаления"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H "Смета успешно удалена"
// @Failure 400 {object} gin.H "Неверный запрос или неверный ID сметы"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Нет прав доступа"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/estimates/{id} [delete]
func (c *EstimateController) DeleteEstimate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	estimateID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат Estimate ID"})
		return
	}

	estimate := &models.Estimate{}
	estimate.ID = uint(estimateID)

	if err := c.estimateService.DeleteEstimate(estimate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estimate deleted successfully"})
}

// GetEstimateByID
// @Summary Получить смету по ID
// @Description Возвращает детальную информацию о смете по ID. Доступно всем авторизованным пользователям.
// @Tags Estimates
// @Produce json
// @Param id path integer true "ID сметы для получения"
// @Security ApiKeyAuth
// @Success 200 {object} models.Estimate "Информация о смете"
// @Failure 400 {object} gin.H "Неверный запрос или неверный ID сметы"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 404 {object} gin.H "Смета не найдена"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/estimates/{id} [get]
func (c *EstimateController) GetEstimateByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	estimateID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат Estimate ID"})
		return
	}

	estimate, err := c.estimateService.GetEstimateByID(estimateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if estimate == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Смета не найдена"})
		return
	}

	ctx.JSON(http.StatusOK, estimate)
}

type TestController struct{}

func NewTestController() *TestController {
	return &TestController{}
}

// GetTestEndpoint
// @Summary Тестовый endpoint
// @Description Возвращает тестовое сообщение.
// @Tags Test
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /test [get]
func (c *TestController) GetTestEndpoint(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Test endpoint is working!")
}

// GetEstimateByCompany
// @Summary Получить сметы компании
// @Description Возвращает список смет для указанной компании. Доступно всем авторизованным пользователям.
// @Tags Estimates
// @Produce json
// @Param company_id query integer true "ID компании, сметы которой нужно получить"
// @Security ApiKeyAuth
// @Success 200 {array} models.Estimate "Список смет компании"
// @Failure 400 {object} gin.H "Неверный запрос или неверный Company ID"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/estimates/company [get]
func (c *EstimateController) GetEstimateByCompany(ctx *gin.Context) {
	companyIDStr := ctx.Query("company_id")
	if companyIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company ID обязателен"})
		return
	}

	companyID, err := strconv.ParseUint(companyIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат Company ID"})
		return
	}

	estimates, err := c.estimateService.GetEstimatesByCompanyID(uint(companyID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if estimates == nil { //  estimates может быть nil или пустой срез. nil - ошибка, пустой срез - нет смет
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при получении смет"})
		return
	}

	if len(estimates) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "Сметы для данной компании не найдены", "estimates": []models.Estimate{}}) // возвращаю 200 OK и пустой массив, если нет смет
		return
	}

	ctx.JSON(http.StatusOK, estimates)
}
