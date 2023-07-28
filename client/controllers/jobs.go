package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc-jobs/proto"
)

const (
	port = ":8080"
)

func GetJobs(c *gin.Context) {


	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS!"})
}

func PostJobs(c *gin.Context) {
	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewJobServiceClient(conn)

	job := &pb.Jobs{
		Id: "1",
		Title: "Job Title",
	} 

	CallJobBiStream(client, job)


	c.JSON(http.StatusOK, gin.H{"message": "SUCCESS!"})
}

func CallJobBiStream(client pb.JobServiceClient, job *pb.Jobs) {
	log.Printf("Bidirectional Streaming started")
	stream, err := client.JobsBiStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send products: %v", err)
	}
	
	waitc := make(chan struct{})

	go func() {		
			message, err := stream.Recv()
			if err != nil {
				log.Fatalf("Error while streaming %v", err)
			}
			log.Println(message)
		close(waitc)
	}()

	message := &pb.Jobs{
		Id:    	job.Id,
		Title:  job.Title,
	}

	req := &pb.JobsRequest{
		Message: message,
	}
	
	if err := stream.Send(req); err != nil {
		log.Fatalf("Error while sending %v", err)
	}	

	stream.CloseSend()
	<-waitc
	log.Printf("Bidirectional Streaming finished")
}