package controllers

import (
	"github.com/ProgrammerPeasant/order-control/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const WRONG_ID = "Неверный ID компании"

type CompanyController struct {
	companyService services.CompanyService
}

func NewCompanyController(cs services.CompanyService) *CompanyController {
	return &CompanyController{companyService: cs}
}

// CreateCompany
// @Summary Создать новую компанию
// @Description Создает новую компанию на основе переданных данных. Доступно только администраторам.
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body CreateCompanyRequest true "Данные компании для создания (включая logo_url и цвета)"
// @Security ApiKeyAuth
// @Success 200 {object} models.Company "Компания успешно создана"
// @Failure 400 {object} gin.H "Невалидные данные"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Нет прав доступа"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies [post]
func (c *CompanyController) CreateCompany(ctx *gin.Context) {
	var req CreateCompanyRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные данные: " + err.Error()})
		return
	}

	company, err := c.companyService.Create(req.Name, req.Description, req.Address, req.LogoURL, req.ColorPrimary, req.ColorSecondary, req.ColorAccent)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

// CreateCompanyRequest represents the request body for creating a company with logo and design colors.
type CreateCompanyRequest struct {
	Name           string `json:"name"`
	Description    string `json:"desc"`
	Address        string `json:"address"`
	LogoURL        string `json:"logo_url"`
	ColorPrimary   string `json:"color_primary"`
	ColorSecondary string `json:"color_secondary"`
	ColorAccent    string `json:"color_accent"`
}

// GetCompany
// @Summary Получить информацию о компании по ID
// @Description Возвращает детальную информацию о компании по указанному ID, включая logo_url и цвета. Доступно всем авторизованным пользователям.
// @Tags Companies
// @Produce json
// @Param id path integer true "ID компании"
// @Security ApiKeyAuth
// @Success 200 {object} models.Company "Информация о компании"
// @Failure 400 {object} gin.H "Неверный ID компании"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 404 {object} gin.H "Компания не найдена"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies/{id} [get]
func (c *CompanyController) GetCompany(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": WRONG_ID})
		return
	}

	company, err := c.companyService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена"})
		return
	}

	ctx.JSON(http.StatusOK, company) // Этот метод уже возвращает все поля модели Company
}

// UpdateCompany
// @Summary Обновить информацию о компании
// @Description Обновляет информацию о существующей компании по указанному ID, включая logo_url и цвета. Доступно только администраторам.
// @Tags Companies
// @Accept json
// @Produce json
// @Param id path integer true "ID компании для обновления"
// @Param request body UpdateCompanyRequest true "Новые данные компании (включая logo_url и цвета)"
// @Security ApiKeyAuth
// @Success 200 {object} models.Company "Информация о компании успешно обновлена"
// @Failure 400 {object} gin.H "Неверный ID компании или невалидные данные"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Нет прав доступа"
// @Failure 404 {object} gin.H "Компания не найдена"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies/{id} [put]
func (c *CompanyController) UpdateCompany(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": WRONG_ID})
		return
	}

	var req UpdateCompanyRequest
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные данные: " + err.Error()})
		return
	}

	company, err := c.companyService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена"})
		return
	}

	company.Name = req.Name
	company.Description = req.Desc
	company.Address = req.Address
	company.LogoURL = req.LogoURL
	company.ColorPrimary = req.ColorPrimary
	company.ColorSecondary = req.ColorSecondary
	company.ColorAccent = req.ColorAccent

	if err := c.companyService.Update(company); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

// UpdateCompanyRequest represents the request body for updating a company with logo and design colors.
type UpdateCompanyRequest struct {
	Name           string `json:"name"`
	Desc           string `json:"desc"`
	Address        string `json:"address"`
	LogoURL        string `json:"logo_url"`
	ColorPrimary   string `json:"color_primary"`
	ColorSecondary string `json:"color_secondary"`
	ColorAccent    string `json:"color_accent"`
}

// DeleteCompany
// @Summary Удалить компанию
// @Description Удаляет компанию по указанному ID. Доступно только администраторам.
// @Tags Companies
// @Param id path integer true "ID компании для удаления"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H{message=string} "Компания успешно удалена"
// @Failure 400 {object} gin.H "Неверный ID компании"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Нет прав доступа"
// @Failure 404 {object} gin.H "Компания не найдена или уже удалена"
// @Failure 500 {object} gin.H "Ошибка сервера"
// @Router /v1/companies/{id} [delete]
func (c *CompanyController) DeleteCompany(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": WRONG_ID})
		return
	}

	if err := c.companyService.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена или уже удалена"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Компания удалена"})
}
