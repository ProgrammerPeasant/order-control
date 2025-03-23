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

// Register
// @Summary Зарегистрировать нового пользователя
// @Description Регистрирует нового пользователя в системе. Доступно всем.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body controllers.RegisterRequest true "Данные для регистрации пользователя"
// @Success 200 {object} gin.H{message=string} "Регистрация прошла успешно"
// @Failure 400 {object} gin.H "Невалидные данные запроса"
// @Failure 500 {object} gin.H "Ошибка сервера при регистрации"
// @Router /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Ошибка чтения тела запроса:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		return
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	log.Printf("Register запрос: %s", string(bodyBytes))

	var request RegisterRequest
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

// Login
// @Summary Авторизовать пользователя
// @Description Авторизует пользователя и возвращает JWT токен. Доступно всем.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body controllers.LoginRequest true "Учетные данные пользователя"
// @Success 200 {object} controllers.LoginResponse "Успешная авторизация"
// @Failure 400 {object} gin.H "Невалидные данные запроса"
// @Failure 401 {object} gin.H "Неверные учётные данные"
// @Failure 500 {object} gin.H "Ошибка сервера при авторизации"
// @Router /login [post]
func (c *AuthController) Login(ctx *gin.Context) {
	var request LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		return
	}

	user, token, err := c.userService.Login(request.Username, request.Password)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учётные данные"})
		return
	}

	ctx.JSON(http.StatusOK, LoginResponse{
		Token:    token,
		Username: user.Username,
		Role:     user.Role,
		UserID:   user.ID,
	})
}

// RegisterRequest represents the request body for user registration.
type RegisterRequest struct {
	Username  string `json:"username"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CompanyID string `json:"company_id"`
}

// LoginRequest represents the request body for user login.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the response body after successful login.
type LoginResponse struct {
	Token    string `json:"token"`
	Username string `json:"username"`
	Role     string `json:"role"`
	UserID   uint   `json:"userId"`
}
