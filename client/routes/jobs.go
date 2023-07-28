package routes

import (
	"grpc-jobs/client/controllers"

	"github.com/gin-gonic/gin"
)

func ProductRoutes(router *gin.Engine) {
	router.GET("/jobs", controllers.GetJobs)
	router.POST("/job", controllers.PostJobs)
}