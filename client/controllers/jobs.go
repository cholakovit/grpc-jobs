package controllers

import (
	"net/http"

	clientgrpc "grpc-jobs/client/clientGrpc"
	"grpc-jobs/client/constants"
	"grpc-jobs/server/model"

	"github.com/gin-gonic/gin"
)

var (
	job		*model.Job
)

func GetJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{constants.MESSAGE: constants.SUCCESS})
}

func JobForm(c *gin.Context) {
	c.HTML(http.StatusOK, "job.html", nil)
}

func PostJob(c *gin.Context) {

	postedJob := model.Job{
		Title: c.PostForm("title"),
		Description: c.PostForm("description"),
		Company: c.PostForm("company"),
		Location: c.PostForm("location"),
		Employment_Type: c.PostForm("employment_type"),
		Salary: c.PostForm("salary"),
		Requirements: c.PostForm("requirements"),
		Responsibilities: c.PostForm("responsibilities"),
		Contact_Information: c.PostForm("contact_information"),
		Application_Process: c.PostForm("application_process"),
	}

	clientgrpc.ConnectServer(&postedJob)

	// Redirect back to the JobForm handler
	c.Redirect(http.StatusSeeOther, "/job")
}