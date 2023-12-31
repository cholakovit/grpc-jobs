package clientgrpc

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc-jobs/proto"
)

// Define a channel to pass the received message to the controller
var MessageChan = make(chan string, 1) // Use a buffer size of 1

func CallCreateJob() {

}

func CallJobBiStream(client pb.JobServiceClient, job *pb.Job) {
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
		log.Println("Comming from the server", message.Message)

		// Send the received message to the channel (non-blocking)
		select {
		case MessageChan <- message.Message:
				// Message sent successfully
		default:
				// Channel buffer is full, handle the case if needed
				log.Println("Channel buffer is full, dropping the message.")
		}

		close(waitc)
	}()

	message := &pb.Job{
		Id:                 job.Id,
		Title:              job.Title,
		Description:        job.Description,
		Company:            job.Company,
		Location:           job.Location,
		Salary:             job.Salary,
		Requirements:       job.Requirements,
		Responsibilities:   job.Responsibilities,
		ContactInformation: job.ContactInformation,
		ApplicationProcess: job.ApplicationProcess,
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

func ReceiveJobs(client pb.JobServiceClient) (*pb.JobListResponse, error) {
	fmt.Println("Receive Jobs client")
	ctx, cancel := context.WithTimeout(context.Background(), 55*time.Second)
	defer cancel()

	res, err := client.ReturnJobList(ctx, &pb.NoParam{})
	if err != nil {
		log.Fatalf("Could not receive jobs: %v", err)
	}

	return res, nil

	//fmt.Println("Receive Jobs client res: ", res)

	// for _, job := range res.Jobs {
	// 	fmt.Printf("Job ID: %s\n", job.Id)
	// 	fmt.Printf("Title: %s\n", job.Title)
	// 	fmt.Printf("Description: %s\n", job.Description)
	// 	// ... continue accessing other job fields
	// 	fmt.Println("====================")
	// }

	//log.Printf("%s", res.Message)
	//log.Printf("%s", res.Jobs)
}