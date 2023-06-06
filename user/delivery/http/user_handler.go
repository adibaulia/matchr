package http

import (
	"net/http"

	"github.com/adibaulia/matchr/domain"
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
			Gender:      userRegister.Gender,
		},
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			StatusCode: http.StatusBadRequest,
			Status:     "error",
			Message:    err.Error(),
		})
	}

	c.JSON(200, nil)
}
