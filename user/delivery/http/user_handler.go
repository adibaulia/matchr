package http

import (
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
	matchr.POST("", handler.enableWallet)

}
