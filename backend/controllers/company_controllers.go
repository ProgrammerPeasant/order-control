package controllers

import (
	"github.com/ProgrammerPeasant/order-control/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CompanyController struct {
	companyService services.CompanyService
}

func NewCompanyController(cs services.CompanyService) *CompanyController {
	return &CompanyController{companyService: cs}
}

func (c *CompanyController) CreateCompany(ctx *gin.Context) {
	var req struct {
		Name    string `json:"name"`
		Desc    string `json:"desc"`
		Address string `json:"address"`
	}

	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные данные"})
		return
	}

	company, err := c.companyService.Create(req.Name, req.Desc, req.Address)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func (c *CompanyController) GetCompany(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компании"})
		return
	}

	company, err := c.companyService.GetByID(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена"})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func (c *CompanyController) UpdateCompany(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компании"})
		return
	}

	var req struct {
		Name    string `json:"name"`
		Desc    string `json:"desc"`
		Address string `json:"address"`
	}
	if err := ctx.BindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидные данные"})
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

	if err := c.companyService.Update(company); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, company)
}

func (c *CompanyController) DeleteCompany(ctx *gin.Context) {
	idParam := ctx.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный ID компании"})
		return
	}

	if err := c.companyService.Delete(uint(id)); err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Компания не найдена или уже удалена"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Компания удалена"})
}
