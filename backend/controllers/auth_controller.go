package controllers

import (
	"bytes"
	"github.com/ProgrammerPeasant/order-control/services"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const invReq = "Невалидный запрос"

type AuthController struct {
	userService services.UserService
}

func NewAuthController(us services.UserService) *AuthController {
	return &AuthController{userService: us}
}

func (c *AuthController) Register(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Ошибка чтения тела запроса:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		return
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	log.Printf("Register запрос: %s", string(bodyBytes))

	var request struct {
		Username  string `json:"username"`
		Email     string `json:"email"`
		Password  string `json:"password"`
		Role      string `json:"role"`
		CompanyID string `json:"company_id"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		log.Println(err)
		return
	}

	companyIDUint, err := strconv.ParseUint(request.CompanyID, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Неверный формат Company ID"})
		return
	}

	err = c.userService.Register(request.Username, request.Email, request.Password, request.Role, uint(companyIDUint))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Регистрация прошла успешно"})
}

func (c *AuthController) Login(ctx *gin.Context) {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		return
	}

	user, token, err := c.userService.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учётные данные"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"token":    token,
		"username": user.Username, // Возвращаю имя пользователя
		"role":     user.Role,     // Возвращаю роль пользователя
		"userId":   user.ID,       // Возвращаю ID пользователя
		// ...
	})
}
