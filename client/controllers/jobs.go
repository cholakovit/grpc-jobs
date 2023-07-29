package controllers

import (
	"net/http"

	clientgrpc "grpc-jobs/client/clientGrpc"
	"grpc-jobs/server/model"

	"github.com/gin-gonic/gin"
)

var (
	job		*model.Job
)

func GetJobs(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS!"})
}

func JobForm(c *gin.Context) {
	c.HTML(http.StatusOK, "job.html", nil)
}

func PostJob(c *gin.Context) {
	title := c.PostForm("title")
	description := c.PostForm("description")
	company := c.PostForm("company")
	location := c.PostForm("location")
	employment_type := c.PostForm("employment_type")
	salary := c.PostForm("salary")
	requirements := c.PostForm("requirements")
	responsibilities := c.PostForm("responsibilities")
	contact_information := c.PostForm("contact_information")
	application_process := c.PostForm("application_process")

	postedJob := model.Job{
		Title: title,
		Description: description,
		Company: company,
		Location: location,
		Employment_Type: employment_type,
		Salary: salary,
		Requirements: requirements,
		Responsibilities: responsibilities,
		Contact_Information: contact_information,
		Application_Process: application_process,
	}

	clientgrpc.ConnectServer(&postedJob)

	// Redirect back to the JobForm handler
	c.Redirect(http.StatusSeeOther, "/job")
}