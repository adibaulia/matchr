package main

import (
	"fmt"

	"github.com/adibaulia/matchr/config"
	"github.com/adibaulia/matchr/user/delivery/http"
	userRepository "github.com/adibaulia/matchr/user/repository/postgre"
	userUsecase "github.com/adibaulia/matchr/user/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	conn := config.GetConnection()

	db, err := conn.PostgreCon.DB()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	r := gin.New()

	userRepo := userRepository.NewUserRepository(conn.PostgreCon)
	userUcase := userUsecase.NewUserUsecase(userRepo)

	http.NewWalletHandler(r, userUcase)

	r.Run(fmt.Sprintf(":%v", config.Conf.ServicePort))
}
