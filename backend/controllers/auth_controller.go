package controllers

import (
	"bytes"
	"github.com/ProgrammerPeasant/order-control/models"
	"github.com/ProgrammerPeasant/order-control/services"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userService services.UserService
}

func NewAuthController(us services.UserService) *AuthController {
	return &AuthController{userService: us}
}

func (c *AuthController) Register(ctx *gin.Context) {
	bodyBytes, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Ошибка чтения тела запроса:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный запрос"})
		return
	}
	// Восстанавливаем тело запроса, так как оно было прочитано
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	log.Printf("Register запрос: %s", string(bodyBytes))

	var request struct {
		Username string      `json:"username"`
		Email    string      `json:"email"`
		Password string      `json:"password"`
		Role     models.Role `json:"role"`
	}

	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный запрос"})
		log.Println(err)
		return
	}

	err = c.userService.Register(request.Username, request.Email, request.Password, request.Role)
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
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Невалидный запрос"})
		return
	}
	token, err := c.userService.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учётные данные"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}
