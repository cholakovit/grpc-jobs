package routes

import (
	"grpc-jobs/client/constants"
	"grpc-jobs/client/controllers"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func JobRoutes(r *gin.Engine) {

	baseDir := filepath.Dir(os.Args[0])
	templatesPath := filepath.Join(baseDir, constants.TEMPLATES, constants.DOTHTML)
	r.LoadHTMLGlob(templatesPath)

	r.Static("/templates", "./templates")

	r.GET("/", controllers.Home)
	r.GET("/jobs", controllers.GetJobs)
	r.GET("/job", controllers.JobForm)
	r.POST("/postJob", controllers.PostJob)
}