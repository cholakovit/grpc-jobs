package controllers

import (
	"net/http"

	clientgrpc "grpc-jobs/client/clientGrpc"
	"grpc-jobs/client/constants"
	"grpc-jobs/client/helper"
	"grpc-jobs/server/model"

	"github.com/gin-gonic/gin"
)

var (
	job		*model.Job
	PORT = helper.LoadEnv("PORT")
	LOCALHOST = helper.LoadEnv("LOCALHOST")
)

type TemplateData struct {
	ReceivedMessage string
}

func renderMenu(c *gin.Context) {

	c.HTML(http.StatusOK, "menu.html", nil)
}

func Home(c *gin.Context) {

	// Render the template with the data
	c.HTML(http.StatusOK, "home.html", nil)
}

func GetJobs(c *gin.Context) {

	client := clientgrpc.ConnectServer()
	clientgrpc.ReceiveJobs(client)

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

	clientgrpc.PostJobServer(&postedJob)

  // Receive the message from the channel (this will block until a message is received)
  receivedMessage := <-clientgrpc.MessageChan

	data := TemplateData{
		ReceivedMessage: receivedMessage,
	}

	// Redirect back to the JobForm handler
	//c.Redirect(http.StatusSeeOther, "/job", data)

	// Render the template with the data
	c.HTML(http.StatusOK, "job.html", data)
}