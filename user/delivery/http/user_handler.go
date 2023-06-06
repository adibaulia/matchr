package http

import (
	"net/http"
	"strings"

	"github.com/adibaulia/matchr/domain"
	"github.com/adibaulia/matchr/pkg/token"
	"github.com/gin-gonic/gin"
)

type (
	userHandler struct {
		userUseCase domain.UserUseCase
	}
)

func NewWalletHandler(r *gin.Engine, us domain.UserUseCase) {
	handler := &userHandler{us}
	r.SetTrustedProxies(nil)

	rg := r.Group("/api/v1")

	matchr := rg.Group("/matchr")
	matchr.POST("/register", handler.register)
	matchr.POST("/login", handler.login)
}

func (h *userHandler) login(c *gin.Context) {
	var userRegister domain.UserRequest
	c.Bind(&userRegister)

	token, err := h.userUseCase.LoginUser(userRegister.Username, userRegister.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	c.JSON(200, domain.Response{
		StatusCode: 200,
		Status:     "success",
		Data: domain.LoginResponse{
			Token: token,
		},
	})
}

func extractToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, domain.Response{
			StatusCode: http.StatusUnauthorized,
			Status:     "errror",
			Message:    "Missing Authorization header",
		})

	}
	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	userID, err := token.ValidateToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			StatusCode: http.StatusUnauthorized,
			Status:     "errror",
			Message:    err.Error(),
		})
	}
	c.Set("userID", userID)
	c.Next()
}

func (h *userHandler) register(c *gin.Context) {
	var userRegister domain.UserRequest
	c.Bind(&userRegister)

	err := h.userUseCase.RegisterUser(domain.User{
		Username: userRegister.Username,
		Email:    userRegister.Email,
		Password: userRegister.Password,
		Profile: domain.Profile{
			Name:        userRegister.Name,
			DateOfBirth: userRegister.DateOfBirth,
			Bio:         userRegister.Bio,
			Gender:      domain.GetGender(userRegister.Gender),
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err.Error(),
		})
		return
	}

	c.JSON(200, nil)
}
