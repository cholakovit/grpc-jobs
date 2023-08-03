package routes

import (
	"grpc-jobs/client/helper"

	"github.com/gin-gonic/gin"
)

var (
	PORT = helper.LoadEnv("PORT")
	LOCALHOST = helper.LoadEnv("LOCALHOST")
)

func InitRoutes() {
	router := gin.Default()
	JobRoutes(router)

	router.Run(LOCALHOST + PORT)
}
