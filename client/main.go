package main

import (
	"grpc-jobs/client/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)	
	//gin.SetMode(gin.DebugMode)	

	routes.InitRoutes()
}