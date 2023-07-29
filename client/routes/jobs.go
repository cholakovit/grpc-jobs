package routes

import (

	"grpc-jobs/client/controllers"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func JobRoutes(router *gin.Engine) {

	baseDir := filepath.Dir(os.Args[0])
	templatesPath := filepath.Join(baseDir, "templates", "*.html")
	router.LoadHTMLGlob(templatesPath)

	router.GET("/jobs", controllers.GetJobs)
	router.GET("/job", controllers.JobForm)
	router.POST("/postJob", controllers.PostJob)
}