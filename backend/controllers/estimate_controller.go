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

func (c *EstimateController) CreateEstimate(ctx *gin.Context) {
	var estimate models.Estimate
	if err := ctx.ShouldBindJSON(&estimate); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	currentUser := ctx.MustGet("user").(*models.User)
	estimate.CreatedByID = currentUser.ID

	if err := c.estimateService.CreateEstimate(&estimate); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, estimate)
}

func (c *EstimateController) UpdateEstimate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	estimateID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid estimate ID"})
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

func (c *EstimateController) DeleteEstimate(ctx *gin.Context) {
	idStr := ctx.Param("id")
	estimateID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid estimate ID"})
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

func (c *EstimateController) GetEstimateByID(ctx *gin.Context) {
	idStr := ctx.Param("id")
	estimateID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid estimate ID"})
		return
	}

	estimate, err := c.estimateService.GetEstimateByID(estimateID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if estimate == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"message": "Estimate not found"})
		return
	}

	ctx.JSON(http.StatusOK, estimate)
}

func (c *EstimateController) GetEstimateByCompany(ctx *gin.Context) {
	companyIDStr := ctx.Query("company_id")
	if companyIDStr == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Company ID is required"})
		return
	}

	companyID, err := strconv.ParseUint(companyIDStr, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Company ID"})
		return
	}

	estimates, err := c.estimateService.GetEstimatesByCompanyID(uint(companyID))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if estimates == nil { //  estimates может быть nil или пустой срез. nil - ошибка, пустой срез - нет смет
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error fetching estimates"})
		return
	}

	if len(estimates) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No estimates found for this company", "estimates": []models.Estimate{}}) // возвращаю 200 OK и пустой массив, если нет смет
		return
	}

	ctx.JSON(http.StatusOK, estimates)
}
