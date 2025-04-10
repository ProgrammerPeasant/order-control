package controllers

import (
	"bytes"
	"github.com/ProgrammerPeasant/order-control/services"
	"github.com/ProgrammerPeasant/order-control/utils"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

const invReq = "Невалидный запрос"

type AuthController struct {
	userService services.UserService
	metrics     *utils.Metrics
}

func NewAuthController(us services.UserService, m *utils.Metrics) *AuthController {
	return &AuthController{
		userService: us,
		metrics:     m,
	}
}

// Register
// @Summary Зарегистрировать нового пользователя (обычная регистрация)
// @Description Регистрирует нового пользователя в системе. Доступно всем. Пользователь отправляет запрос на присоединение к компании.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body controllers.StandardRegisterRequest true "Данные для регистрации пользователя (логин, email, пароль, CompanyID)"
// @Success 200 {object} gin.H{message=string} "Регистрация прошла успешно. Запрос на присоединение к компании отправлен на одобрение."
// @Failure 400 {object} gin.H "Невалидные данные запроса"
// @Failure 500 {object} gin.H "Ошибка сервера при регистрации"
// @Router /register [post]
func (c *AuthController) Register(ctx *gin.Context) {
	startTime := time.Now()

	var request StandardRegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		c.metrics.RegisterError("validation_error", "Invalid register request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		log.Println(err)
		return
	}

	user, err := c.userService.Register(request.Username, request.Email, request.Password, "USER", 0) // CompanyID пока 0
	if err != nil {
		c.metrics.RegisterError("registration_error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.userService.CreateJoinRequest(user.ID, request.CompanyID, request.Email)
	if err != nil {
		c.metrics.RegisterError("join_request_error", err.Error())
		// Обработка ошибки (например, удаление пользователя, если запрос не создался)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка при создании запроса на присоединение к компании"})
		return
	}

	// Регистрируем успешную операцию регистрации
	c.metrics.TotalRequests.WithLabelValues("POST", "/register", "200").Inc()

	// Регистрируем время операции
	elapsed := time.Since(startTime).Seconds()
	c.metrics.AuthOperationDuration.WithLabelValues("register").Observe(elapsed)

	// Увеличиваем счетчик успешных регистраций
	c.metrics.UserRegistrations.WithLabelValues("standard").Inc()

	ctx.JSON(http.StatusOK, gin.H{"message": "Регистрация прошла успешно. Запрос на присоединение к компании отправлен на одобрение."})
}

// AdminRegister
// @Summary Зарегистрировать нового пользователя (только для администраторов)
// @Description Регистрирует нового пользователя в системе с указанием роли и компании. Требуется роль администратора.
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body controllers.AdminRegisterRequest true "Данные для регистрации пользователя (включая роль и CompanyID)"
// @Security ApiKeyAuth
// @Success 200 {object} gin.H{message=string} "Регистрация прошла успешно"
// @Failure 400 {object} gin.H "Невалидные данные запроса"
// @Failure 401 {object} gin.H "Не авторизован"
// @Failure 403 {object} gin.H "Доступ запрещен. Требуется роль администратора."
// @Failure 500 {object} gin.H "Ошибка сервера при регистрации"
// @Router /admin/register [post]
func (c *AuthController) AdminRegister(ctx *gin.Context) {
	startTime := time.Now()

	bodyBytes, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		c.metrics.RegisterError("request_read_error", err.Error())
		log.Println("Ошибка чтения тела запроса:", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		return
	}

	ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	log.Printf("AdminRegister запрос: %s", string(bodyBytes))

	var request AdminRegisterRequest
	if err := ctx.BindJSON(&request); err != nil {
		c.metrics.RegisterError("validation_error", "Invalid admin register request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		log.Println(err)
		return
	}

	_, err = c.userService.Register(request.Username, request.Email, request.Password, request.Role, request.CompanyID)
	if err != nil {
		c.metrics.RegisterError("admin_registration_error", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Регистрируем успешную операцию регистрации админом
	c.metrics.TotalRequests.WithLabelValues("POST", "/admin/register", "200").Inc()

	// Регистрируем время операции
	elapsed := time.Since(startTime).Seconds()
	c.metrics.AuthOperationDuration.WithLabelValues("admin_register").Observe(elapsed)

	// Увеличиваем счетчик успешных регистраций
	c.metrics.UserRegistrations.WithLabelValues("admin").Inc()

	// Отслеживаем регистрации по ролям
	c.metrics.UserRoleRegistrations.WithLabelValues(request.Role).Inc()

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
	startTime := time.Now()

	var request LoginRequest
	if err := ctx.BindJSON(&request); err != nil {
		c.metrics.RegisterError("validation_error", "Invalid login request format")
		ctx.JSON(http.StatusBadRequest, gin.H{"error": invReq})
		return
	}

	user, token, err := c.userService.Login(request.Username, request.Password)
	if err != nil {
		c.metrics.RegisterError("authentication_error", "Invalid credentials")
		c.metrics.AuthFailures.WithLabelValues(request.Username).Inc()
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Неверные учётные данные"})
		return
	}

	// Регистрируем успешную операцию входа
	c.metrics.TotalRequests.WithLabelValues("POST", "/login", "200").Inc()

	// Регистрируем время операции
	elapsed := time.Since(startTime).Seconds()
	c.metrics.AuthOperationDuration.WithLabelValues("login").Observe(elapsed)

	// Увеличиваем счетчик успешных входов
	c.metrics.UserLogins.WithLabelValues(user.Role).Inc()

	// Отслеживаем активных пользователей
	c.metrics.ActiveUsers.Inc()

	ctx.JSON(http.StatusOK, LoginResponse{
		Token:     token,
		Username:  user.Username,
		Role:      user.Role,
		UserID:    user.ID,
		CompanyID: user.CompanyID,
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

// LoginResponse represents the response body after successful login.
type LoginResponse struct {
	Token     string `json:"token"`
	Username  string `json:"username"`
	Role      string `json:"role"`
	UserID    uint   `json:"userId"`
	CompanyID uint   `json:"companyId"`
}
