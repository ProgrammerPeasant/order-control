package middlewares

import (
	"encoding/json"
	"github.com/ProgrammerPeasant/order-control/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ProgrammerPeasant/order-control/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthMiddlewareTestSuite struct {
	suite.Suite
	router *gin.Engine
	claims utils.Claims
}

func (s *AuthMiddlewareTestSuite) SetupSuite() {
	gin.SetMode(gin.TestMode)

	s.router = gin.New()

	s.router.GET("/protected", AuthMiddleware(), func(c *gin.Context) {
		userID, _ := c.Get("userID")
		role, _ := c.Get("role")
		companyID, _ := c.Get("companyID")

		c.JSON(http.StatusOK, gin.H{
			"userID":    userID,
			"role":      role,
			"companyID": companyID,
		})
	})

	s.claims = utils.Claims{
		UserID:    123,
		Role:      "admin",
		CompanyID: 456,
	}
}

func (s *AuthMiddlewareTestSuite) TestNoAuthHeader() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
}

func (s *AuthMiddlewareTestSuite) TestInvalidHeaderFormat() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "InvalidToken123")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)

	req, _ = http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer Token Extra")
	w = httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
}

func (s *AuthMiddlewareTestSuite) TestInvalidToken() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer InvalidTokenString")
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
}

func (s *AuthMiddlewareTestSuite) TestValidToken() {
	user := &models.User{
		Role:      s.claims.Role,
		CompanyID: s.claims.CompanyID,
		Email:     "test@example.com",
	}

	token, err := utils.GenerateJWT(user)
	assert.NoError(s.T(), err)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)

	var response map[string]interface{}
	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(s.T(), err)

	assert.Equal(s.T(), float64(0), response["userID"])
	assert.Equal(s.T(), s.claims.Role, response["role"])
	assert.Equal(s.T(), float64(s.claims.CompanyID), response["companyID"])
}

func TestAuthMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareTestSuite))
}
